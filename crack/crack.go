package crack

import (
	"github.com/bernardoaraujor/corinda/train"
	"crypto/sha256"
	"encoding/gob"
	"crypto/sha1"
	"runtime"
	"hash"
	"sync"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"encoding/hex"
)

const Rockyou = "rockyou"
const Linkedin = "linkedin"
const Antipublic = "antipublic"

const SHA1 = "SHA1"
const SHA256 = "SHA256"

type targetsMap struct{
	sync.RWMutex
	targets map[string]string
}

type password struct {
	pass string
	hash []byte
}

type Crack struct {
	alg        string
	trainMaps  train.Maps
	targetsMap targetsMap
}

// crack session
func (crack Crack) Crack(){
	composites := crack.trainMaps.CompositeMap
	var wg sync.WaitGroup

	// initialize channels
	guesses := make([]chan string, 0)
	guessesPerBatch := make([]int, 0)
	for _, cm := range composites{
		guesses = append(guesses, cm.Guess())
		guessesPerBatch = append(guessesPerBatch, int(cm.Entropy))
	}

	batches := crack.batch(guesses, guessesPerBatch)

	// TODO: many parallell searchers?
	wg.Add(1)
	results := crack.searcher(batches)
	go save(results, wg)

	fmt.Println("Cracking...")
	wg.Wait()
}

// Constructor
func NewCrack(list string, alg string) Crack{
	var crack Crack

	crack.alg = alg

	err := load("maps/"+ list +"Maps.gob", &crack.trainMaps)
	check(err)

	f, err := ioutil.ReadFile("targets/rockyouSHA1.csv")
	check(err)
	targets := strings.Split(string(f), "\n")

	fmt.Println("Loading target list...")
	crack.targetsMap.targets = make(map[string]string)
	for _, hash := range targets{
		crack.targetsMap.targets[hash] = hash
	}

	return crack
}

// Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
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

func save(in chan password, wg sync.WaitGroup){
	defer wg.Done()

	resultFile, err := os.Create("results/testResults.csv")
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