package crack

import (
	"crypto/sha256"
	"crypto/sha1"
	"runtime"
	"hash"
	"sync"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/hex"
	"github.com/bernardoaraujor/corinda/composite"
	"github.com/bernardoaraujor/corinda/elementary"
	"encoding/json"
	"compress/gzip"
	"encoding/csv"
	"math"
)

const Rockyou = "rockyou"
const Linkedin = "linkedin"
const Antipublic = "antipublic"

const SHA1 = "sha1"
const SHA256 = "sha256"

type targetsMap struct{
	sync.RWMutex
	targets map[string]string
}

type password struct {
	pass string
	hash []byte
}

type Crack struct {
	alg          string
	composites   map[string]*composite.Model
	elementaries map[string]*elementary.Model
	targetsMap   targetsMap
	targetName	 string
	trainedName	 string
}

// crack session
func (crack Crack) Crack(){
	composites := crack.composites
	elementaries := crack.elementaries

	var wg sync.WaitGroup

	// initialize channels
	fmt.Println("Initializing Guess Channels")

	guesses := make([]chan string, 0)
	guessesPerBatch := make([]int, 0)
	for _, cm := range composites{
		tokenLists := make([][]string, 0)
		for _, elementaryName := range cm.Models{
			elementary := elementaries[elementaryName]

			tokenLists = append(tokenLists, elementary.SortedTokens())
		}

		guesses = append(guesses, cm.Guess(tokenLists))

		// gpb = k*p*10^E
		k := 1.0
		gpb := int(k * cm.Prob*math.Pow(10, cm.Entropy))
		guessesPerBatch = append(guessesPerBatch, gpb)
	}

	batches := crack.batch(guesses, guessesPerBatch)

	// TODO: many parallell searchers?
	wg.Add(1)
	results := crack.searcher(batches)
	go save(crack.trainedName, crack.targetName, crack.alg, results, wg)

	fmt.Println("Cracking...")
	wg.Wait()
}

// Constructor
func NewCrack(trained string, target string, alg string) Crack{
	var crack Crack

	crack.targetName = target
	crack.trainedName = trained
	crack.alg = alg

	f, err := os.Open("maps/"+ trained + "Elementaries.json.gz")
	check(err)

	gr, err := gzip.NewReader(f)
	check(err)

	raw, err := ioutil.ReadAll(gr)
	check(err)
	var elementaries []*elementary.Model
	err = json.Unmarshal(raw, &elementaries)
	check(err)
	crack.elementaries = make(map[string]*elementary.Model)

	for _, em := range elementaries{
		crack.elementaries[em.Name] = em
	}

	f, err = os.Open("maps/"+ trained +"Composites.json.gz")
	check(err)

	gr, err = gzip.NewReader(f)
	check(err)

	raw, err = ioutil.ReadAll(gr)
	check(err)
	var composites []*composite.Model
	err = json.Unmarshal(raw, &composites)

	crack.composites = make(map[string]*composite.Model)

	for _, cm := range composites{
		crack.composites[cm.Name] = cm
	}

	f, err = os.Open("targets/" + alg + "/" + target + ".csv.gz")
	check(err)
	defer f.Close()

	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cr := csv.NewReader(gr)

	fmt.Println("Loading target list...")
	crack.targetsMap.targets = make(map[string]string)
	for records, err := cr.Read(); records != nil; records, err = cr.Read(){
		check(err)

		hash := records[0]
		crack.targetsMap.targets[hash] = hash
	}

	return crack
}

// Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

// returns channel with hashes in string format, and iterates n times retrieving password guesses over in
// digest implements the fan in patterns
func (crack Crack) digest(in chan string, n int) chan password {
	out := make(chan password)

	go func(n int, out chan password) {
		defer close(out)
		for i := 0; i < n; i++ {
			// reads in channel
			guess := <-in

			// temporary conditional... avoiding weird bug that receives empty guess
			if guess != ""{
				// digest
				var hasher hash.Hash
				switch crack.alg {
				case SHA1:
					hasher = sha1.New()
				case SHA256:
					hasher = sha256.New()
				}
				hasher.Write([]byte(guess))
				digest := hasher.Sum(nil)

				// spits out digest
				out <- password{guess, digest}
			}
		}
	}(n, out)

	return out
}

// merge the flux from channels cs into out
func fanIn(cs []chan password) chan password {
	var wg sync.WaitGroup

	out := make(chan password)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan password) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// prepares wait group for the number of input channels
	wg.Add(len(cs))

	// start goroutines
	for _, c := range cs {
		go output(c)
	}

	// start goroutine that closes output channel when all input channels have been closed
	go func() {
		wg.Wait()
		close(out)
	}()

	// returns output channel
	return out
}

// generate a batch of hashes from the input guess channels
func (crack Crack) batch(guessChans []chan string, ns []int) chan []password {
	out := make(chan []password)

	go func(guessChans []chan string, ns []int){
		for {
			// generates array of channels for digesting, to be used as inputs to fanIn
			passwords := make([]chan password, 0)
			for i, guessChan := range guessChans{
				n := ns[i]
				passwords = append(passwords, crack.digest(guessChan, n))
			}

			// generate fan in channel
			fanIn := fanIn(passwords)

			// initialize batch of hashes
			batch := make([]password, 0)

			// drain fanIn channel
			for password := range fanIn {
				batch = append(batch, password)
			}

			out <- batch
		}
	}(guessChans, ns)

	return out
}

func (crack *Crack) searcher(in chan []password) chan password {
	out := make(chan password)

	go func(){
		for{
			batch := <- in

			for _, password := range batch{
				//pass := ph.pass
				hash := hex.EncodeToString(password.hash)

				if _, ok := crack.targetsMap.targets[hash]; ok{
					out <- password

					crack.targetsMap.Lock()
					delete(crack.targetsMap.targets, hash)
					crack.targetsMap.Unlock()
				}
			}
		}
	}()

	return out
}

func save(trained string, target string, alg string, in chan password, wg sync.WaitGroup){
	defer wg.Done()

	resultFile, err := os.Create("results/" + trained + "_" + target + "_" + alg +".csv")
	check(err)
	defer resultFile.Close()

	for ph := range in{
		pass := ph.pass
		hash := ph.hash

		line := pass + "," + hex.EncodeToString(hash)
		//fmt.Println(pass)
		fmt.Fprintln(resultFile, line)
	}
}