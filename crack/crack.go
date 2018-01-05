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
	"time"
	"strconv"
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
	durationH	 float64
}

// crack session
func (crack Crack) Crack(){
	composites := crack.composites
	elementaries := crack.elementaries

	var wg sync.WaitGroup

	// initialize channels
	fmt.Println("Initializing Guess Channels")

	guesses := make([]chan string, 0)
	nGuesses := make([]int, 0)
	for _, cm := range composites{
		tokenLists := make([][]string, 0)
		for _, elementaryName := range cm.Models{
			elementary := elementaries[elementaryName]

			tokenLists = append(tokenLists, elementary.SortedTokens())
		}

		guesses = append(guesses, cm.Guess(tokenLists))

		// n = k*p*10^E
		k := 100.0
		n := int(k * cm.Prob * math.Pow(10, cm.Entropy))
		nGuesses = append(nGuesses, n)
	}

	fmt.Println(nGuesses)
	guessLoop := crack.guessLoop(guesses, nGuesses)

	passwords := crack.digest(guessLoop)

	wg.Add(1)
	results := crack.searcher(passwords)
	count := 0

	go monitor(wg, crack.durationH)
	go reporter(&count)
	go save(crack.trainedName, crack.targetName, crack.alg, results, wg, &count)

	fmt.Println("Cracking...")

	wg.Wait()
}

// Constructor
func NewCrack(trained string, target string, alg string, durationH float64) Crack{
	var crack Crack

	crack.targetName = target
	crack.trainedName = trained
	crack.alg = alg
	crack.durationH = durationH

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

// returns channel with hashes in string format
func (crack Crack) digest(in chan string) chan password {
	out := make(chan password)

	go func(out chan password) {
		defer close(out)

		// reads in channel
		for guess := range in {
			// temporary conditional... avoiding weird bug that receives empty guess
			if guess != "" {
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

	}(out)

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

// loops over guess channels generating guesses for digest
func (crack Crack) guessLoop(guessChans []chan string, ns []int) chan string {
	out := make(chan string)

	go func(guessChans []chan string, ns []int){
		for {
			for i, guessChan := range guessChans{
				nGuesses := ns[i]

				for j := 0; j < nGuesses; j++{
					guess := <- guessChan

					out <- guess
				}
			}
		}
	}(guessChans, ns)

	return out
}

func (crack *Crack) searcher(in chan password) chan password {
	out := make(chan password)

	go func(){
		for password := range in{
			//pass := ph.pass
			hash := hex.EncodeToString(password.hash)

			if _, ok := crack.targetsMap.targets[hash]; ok{
				out <- password

				crack.targetsMap.Lock()
				delete(crack.targetsMap.targets, hash)
				crack.targetsMap.Unlock()
			}
		}
	}()

	return out
}

func save(trained string, target string, alg string, in chan password, wg sync.WaitGroup, c *int){
	defer wg.Done()

	resultFile, err := os.Create("results/" + trained + "_" + target + "_" + alg +".csv")
	check(err)
	defer resultFile.Close()

	for ph := range in{
		pass := ph.pass
		hash := ph.hash

		line := pass + "," + hex.EncodeToString(hash)
		*c++

		fmt.Fprintln(resultFile, line)
	}
}

func reporter(c *int){
	lastCount := 0
	// report progress every second
	for {
		time.Sleep(1000 * time.Millisecond)

		speed := *c - lastCount

		fmt.Println("Speed: " + strconv.Itoa(speed) + " P/s; Found Passwords: " + strconv.Itoa(*c))
		lastCount = *c
	}
}

func monitor(wg sync.WaitGroup, durationH float64){
	start := time.Now()

	for{
		since := time.Since(start)
		if since.Hours() > durationH{
			wg.Done()
			os.Exit(0)
		}
	}
}