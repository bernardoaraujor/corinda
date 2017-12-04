package train

import (
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
	"github.com/timob/jnigi"
	"encoding/json"
	"encoding/csv"
	"encoding/gob"
	"strconv"
	"runtime"
	"time"
	"sync"
	"fmt"
	"os"
	"compress/gzip"
)

const passfaultClassPath = "-Djava.class.path=passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"

// this struct is generated after a csv file is processed
type Maps struct {
	ElementaryMap map[string]*elementary.Model
	CompositeMap  map[string]*composite.Model
	N             int
}

func (tmTo *Maps) Merge(tmFrom *Maps){
	tmTo.N += tmFrom.N

	// process ElementaryModelsMaps
	for k, emFrom := range tmFrom.ElementaryMap {
		if emTo, ok := tmTo.ElementaryMap[k]; ok{ //emFrom already in tmTo.ElementaryMap

			//verify every TokenFreq in emFrom
			for _, tf := range emFrom.TokensNfreqs{
				t := tf.Token
				f := tf.Freq

				// is t already in emTo.TokensNfreqs??
				index := 0
				b := false
				for i, tfTo := range emTo.TokensNfreqs{
					if tfTo.Token == t{
						index = i
						b = true
						break
					}
				}

				if b{		// yes, t is in emTo.TokensNfreqs
					emTo.TokensNfreqs[index].Freq += f
				}else{		//no, t is not in emTo.TokensNfreqs
					emTo.TokensNfreqs = append(emTo.TokensNfreqs, elementary.TokenFreq{t, f})
				}

				emTo.Sort()
			}			
		}else{			//emFrom not in tmTo, create new entry
			emFrom.Sort()
			tmTo.ElementaryMap[k] = emFrom
		}
	}

	// process CompositeModels
	for k, cmFrom := range tmFrom.CompositeMap {
		if cmTo, ok := tmTo.CompositeMap[k]; ok{ //cm already in tmTo
			cmTo.UpdateFreq(cmFrom.Freq)
		}else{		//cm not in tmTo, create new entry
			cmTo := cmFrom

			//set em pointers
			for i, em := range cmFrom.Models {
				emName := em.Name
				emTo := tmTo.ElementaryMap[emName]
				cmTo.Models[i] = emTo
			}

			//insert cm
			tmTo.CompositeMap[k] = cmTo
		}
	}

	// update composite models porbabilities
	sum := 0
	for _, cm := range tmTo.CompositeMap{
		sum += cm.Freq
	}

	for _, cm := range tmTo.CompositeMap{
		cm.UpdateProb(sum)
	}
}

// this struct is used right after the csv line is read
// it contains frequency and password
type FreqNpass struct{
	freq int
	pass string
}

// this struct is used after the password has been analyzed
// it contains frequency and a byte slice with the JSON of the analysis
type FreqNresult struct{
	freq int
	result []byte
}

// the types ElModelJSON and CompModelJSON are used only for parsing JSON into ElementaryModel and CompositeModel
// this is done by function DecodeJSON
type ElModelJSON struct{
	ModelName  string `json:"modelName"`
	//Complexity int    `json:"complexity"`
	ModelIndex int    `json:"modelIndex"`
	Token      string `json:"token"`
}

type CompModelJSON struct {
	//Complexity       int `json:"complexity"`
	Models             []ElModelJSON `json:"elementaryModels"`
	CompositeModelName string        `json:"compositeModelName"`
}

// reads lines from csv file and sends them to a buffered channel
// many go routines of ProcessPass will read from this channel
func CsvRead(cr *csv.Reader, nRoutines int) <-chan FreqNpass{

	fpChan := make(chan FreqNpass, nRoutines)
	go func(){
		for records, err := cr.Read(); records != nil; records, err = cr.Read(){
			Check(err)

			freq, err := strconv.Atoi(records[0])
			pass := records[1]
			Check(err)

			fp := FreqNpass{freq, pass}
			fpChan <- fp
		}

		//no more lines, close channel
		close(fpChan)
	}()

	return fpChan
}

// parses the JSON strings returned from Passfault
// data is stored in maps of Composite and Elementary Models
func DecodeJSON(frChan <-chan FreqNresult, done *bool, trainName string){
	compositeModelMap := make(map[string]*composite.Model)
	elementaryModelMap := make(map[string]*elementary.Model)
	nCsvLines := 0

	for { // loop over frChan
		fr, ok := <-frChan
		if ok{ //there are still values to be read
			nCsvLines++
			freq := fr.freq
			result := fr.result

			// parse JSON
			var cmFromJSON CompModelJSON
			json.Unmarshal(result, &cmFromJSON)

			// update elementaryModel map
			for _, emFromJSON := range cmFromJSON.Models {
				if emFromMap, ok := elementaryModelMap[emFromJSON.ModelName]; ok {	//ElementaryModel already in map, only update frequency
					emFromMap.UpdateTokenFreq(freq, emFromJSON.Token)
				}else{	// ElementaryModel not in map, create new instance and insert into the map
					t := emFromJSON.Token
					f := freq
				
					tokensNfreqs := make([]elementary.TokenFreq, 0)
					tokensNfreqs = append(tokensNfreqs, elementary.TokenFreq{t, f})
					//newEM := elementary.Model{emFromJSON.ModelName, emFromJSON.Complexity, 0, tokensNfreqs}
					newEM := elementary.Model{emFromJSON.ModelName, 0, tokensNfreqs}
					elementaryModelMap[emFromJSON.ModelName] = &newEM
				}
			}

			if cmFromMap, ok := compositeModelMap[cmFromJSON.CompositeModelName]; ok{	//CompositeModel already in map, only update frequency
				cmFromMap.UpdateFreq(freq)
			}else{		// CompositeModel not in map, create new instance and insert into the map
				compModelName := cmFromJSON.CompositeModelName
				//complexity := cmFromJSON.Complexity

				// populate array of pointers
				var elementaryModels []*elementary.Model
				for _, emFromJSON := range cmFromJSON.Models {
					emName := emFromJSON.ModelName
					em := elementaryModelMap[emName]
					elementaryModels = append(elementaryModels, em)
				}

				// instantiate new Composite Model
				//cm := composite.Model{compModelName, complexity, freq, 0, elementaryModels}
				cm := composite.Model{compModelName, freq, 0, 0, elementaryModels}

				// add to map
				compositeModelMap[compModelName] = &cm
			}
		}else{ //frChan is closed

			// calculate entropies
			for _, em := range elementaryModelMap{
				em.UpdateEntropy()
			}

			for _, cm := range compositeModelMap{
				cm.UpdateEntropy()
			}

			// calculate composite models probabilities
			sum := 0
			for _, cm := range compositeModelMap{
				sum += cm.Freq
			}

			for _, cm := range compositeModelMap{
				cm.UpdateProb(sum)
			}

			// save file
			trainedMaps := Maps{elementaryModelMap, compositeModelMap, nCsvLines}
			emFile, err := os.Create("maps/" + trainName + "Maps.gob")
			encoder := gob.NewEncoder(emFile)
			err = encoder.Encode(trainedMaps)
			Check(err)
			emFile.Close()

			break
		}
	}

	*done = true
}

// uses Passfault's passwordAnalysis method to process passwords
func PasswordAnalysis(passfaultClassPath string, fpChan <-chan FreqNpass, nRoutines int) ([]<-chan FreqNresult, <-chan bool){
	// start JVM
	jvm, _, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{passfaultClassPath}))
	Check(err)

	frChans := make([]<-chan FreqNresult, nRoutines)

	//this channel is used to keep track of progress
	countChan := make(chan bool, nRoutines)

	//start nRoutines
	for i := 0; i < nRoutines; i++{
		//this channel is used to send results from analysis
		frChan := make(chan FreqNresult, nRoutines)

		frChans[i] = frChan
		go func(){
			// attach this routine to JVM
			env := jvm.AttachCurrentThread()

			// create TextAnalysis JVM object
			obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
			Check(err)

			// loop over inputs from csv file
			for fp := range fpChan{

				// create JVM string with password
				str, err := env.NewObject("java/lang/String", []byte(fp.pass))
				Check(err)

				// call passwordAnalysis on password
				v, err := obj.CallMethod(env, "passwordAnalysis", "java/lang/String", str)
				Check(err)

				// format result from JVM into byte array (probably not the most elegant way!)
				resultJVM, err := v.(*jnigi.ObjectRef).CallMethod(env, "getBytes", jnigi.Byte|jnigi.Array)
				resultString := string(resultJVM.([]byte))
				resultBytes := []byte(resultString)

				// send result to JSON decoder
				frChan <- FreqNresult{fp.freq, resultBytes}

				// signal to counter
				countChan <- true
			}
			close(frChan)
		}()
	}

	return frChans, countChan
}

// fans output of channels from PasswordAnalysis routines into one single channel
func FanIn(frChans []<-chan FreqNresult, nRoutines int) <-chan FreqNresult{
	var wg sync.WaitGroup
	out := make(chan FreqNresult, nRoutines)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan FreqNresult) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(frChans))
	for _, c := range frChans {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// counts how many passwords have already been analyzed
func Counter(c *int, countChan <-chan bool){
	for _ = range countChan{
		*c++
	}
}

// checks for error
func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

func Train(input string, nRoutines int) {
	inputCsvGzPath := "csv/" + input + ".csv.gz"
	//inputCsvPath := "csv/" + input + ".csv"
	f, err := os.Open(inputCsvGzPath)
	//f, err := os.Open(inputCsvPath)
	Check(err)
	defer f.Close()

	gr, err := gzip.NewReader(f)
	Check(err)
	defer gr.Close()

	//cr := csv.NewReader(f)
	cr := csv.NewReader(gr)

	// initialize counter
	count := 0

	//check if pipeline has finished
	var done bool

	// start pipeline
	fpChan := CsvRead(cr, nRoutines)
	frChans, countChan := PasswordAnalysis(passfaultClassPath, fpChan, nRoutines)
	go Counter(&count, countChan)
	fannedFrChans := FanIn(frChans, nRoutines)
	go DecodeJSON(fannedFrChans, &done, input)

	start := time.Now()

	// report progress every second
	for !done{
		since := time.Since(start)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Processed passwords: " + strconv.Itoa(count) + "; Total time: " + since.String())
	}
}



