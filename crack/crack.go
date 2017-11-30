package crack

import (
	"github.com/bernardoaraujor/corinda/train"
	"crypto/sha256"
	//"encoding/csv"
	"encoding/gob"
	"crypto/sha1"
	"runtime"
	"hash"
	"sync"
	"fmt"
	"os"
	"io/ioutil"
	//"encoding/hex"
	"strings"
	"encoding/hex"
)

const Rockyou = "rockyou"
const Linkedin = "linkedin"
const Antipublic = "antipublic"

const SHA1 = "SHA1"
const SHA256 = "SHA256"

type targetMap struct{
	sync.RWMutex
	targets map[string]string
}

type passNhash struct {
	pass string
	hash []byte
}

type Crack struct {
	trainedMaps train.TrainedMaps
	alg         string
	targetMap	targetMap
}

// crack session
func (crack Crack) Crack(){
	fmt.Println(len(crack.targetMap.targets))
	compositeModelsMap := crack.trainedMaps.CompositeModelsMap

	// initialize channels
	guessChans := make([]chan string, 0)
	digestPerBatch := make([]int, 0)
	for _, cm := range compositeModelsMap{
		guessChans = append(guessChans, cm.Guess())
		digestPerBatch = append(digestPerBatch, int(cm.Entropy))
	}

	batchChan := crack.HashBatch(guessChans, digestPerBatch)
	searchChan := make(chan passNhash)
	resultChan := crack.searchTarget(searchChan)
	go saveResults(resultChan)

	fmt.Println("Cracking...")
	for batch := range batchChan{
		for _, ph := range batch{
			searchChan <- ph
		}
	}
}

// Constructor
func NewCrack(list string, alg string) Crack{
	var crack Crack

	crack.alg = alg

	err := load("maps/"+ list +"TrainedMaps.gob", &crack.trainedMaps)
	check(err)

	f, err := ioutil.ReadFile("targets/rockyouSHA1.csv")
	check(err)
	targets := strings.Split(string(f), "\n")

	fmt.Println("Loading target list...")
	crack.targetMap.targets = make(map[string]string)
	for _, hash := range targets{
		crack.targetMap.targets[hash] = hash
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
func (crack Crack) Digest(in chan string, n int) chan passNhash {
	out := make(chan passNhash)

	go func(n int, out chan passNhash) {
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
				out <- passNhash{guess, digest}
			}
		}
	}(n, out)

	return out
}

// merge the flux from channels cs into out
func fanIn(cs []chan passNhash) chan passNhash {
	var wg sync.WaitGroup

	out := make(chan passNhash)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan passNhash) {
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
func (crack Crack) HashBatch(guessChans []chan string, ns []int) chan []passNhash{
	out := make(chan []passNhash)

	go func(guessChans []chan string, ns []int){
		for {
			// generates array of channels for digesting, to be used as inputs to fanIn
			passNhashChan := make([]chan passNhash, 0)
			for i, guessChan := range guessChans{
				n := ns[i]
				passNhashChan = append(passNhashChan, crack.Digest(guessChan, n))
			}

			// generate fan in channel
			fanIn := fanIn(passNhashChan)

			// initialize batch of hashes
			batch := make([]passNhash, 0)

			// drain fanIn channel
			for ph := range fanIn {
				batch = append(batch, ph)
			}

			out <- batch
		}
	}(guessChans, ns)


	return out
}

func (crack *Crack) searchTarget(in chan passNhash) chan passNhash{
	out := make(chan passNhash)

	go func(){
		for{
			ph := <- in
			//pass := ph.pass
			hash := hex.EncodeToString(ph.hash)

			if _, ok := crack.targetMap.targets[hash]; ok{
				out <- ph

				//crack.targetMap.Lock()
				delete(crack.targetMap.targets, hash)
				//crack.targetMap.Unlock()
			}
		}
	}()

	return out
}

func saveResults(in chan passNhash){
	resultFile, err := os.Create("results/testResults.csv")
	check(err)
	defer resultFile.Close()

	for ph := range in{
		pass := ph.pass
		hash := ph.hash

		line := pass + "," + hex.EncodeToString(hash)
		fmt.Fprintln(resultFile, line)
	}
}