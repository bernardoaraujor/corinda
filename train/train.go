package train

import (
	"runtime"
	"fmt"
	"os"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
	"compress/gzip"
	"encoding/csv"
	"strconv"
	"github.com/timob/jnigi"
	"time"
	"encoding/json"
)

const passfaultClassPath = "-Djava.class.path=passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"
const bufSize = 10000000

type input struct {
	freq int
	pass string
}

type inputBatch []input

type result struct {
	freq   int
	result []byte
}

type resultBatch []result

type trainedMaps struct {
	elementaries map[string]*elementary.Model
	composites   map[string]*composite.Model
}

// used only for parsing JSON into elementary.Model
type elementaryJSON struct{
	Name  string `json:"modelName"`
	Index int    `json:"modelIndex"`
	Token      string `json:"token"`
}

// used only for parsing JSON into composite.Model
type compositeJSON struct {
	Models             []elementaryJSON `json:"elementaryModels"`
	CompositeModelName string        `json:"compositeModelName"`
}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

func countLines(list string) int{
	path := "csv/" + list + ".csv.gz"

	f, err := os.Open(path)
	check(err)
	defer f.Close()

	gr, err := gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cr := csv.NewReader(gr)

	//fmt.Println("Counting lines in list...")
	listSize := 0
	for records, err := cr.Read(); records != nil; records, err = cr.Read(){
		check(err)
		listSize++
	}

	return listSize
}

func generator(list string, batchSize int) (int, chan inputBatch){
	listSize := countLines(list)

	out := make(chan inputBatch, bufSize)

	path := "csv/" + list + ".csv.gz"
	f, err := os.Open(path)
	check(err)

	gr, err := gzip.NewReader(f)
	check(err)

	cr := csv.NewReader(gr)

	go func(){
		defer f.Close()
		defer gr.Close()
		end := false

		for{
			ib := make([]input, 0)

			for i := 0; i < batchSize; i++{
				row, _ := cr.Read()

				if row != nil{
					freq, err := strconv.Atoi(row[0])
					pass := row[1]
					check(err)

					input := input{freq, pass}
					ib = append(ib, input)
				}else{	//end of list
					end = true
					break
				}
			}

			out <- ib
			if end{
				close(out)
				return
			}
		}
	}()

	return listSize, out
}

func batchAnalyzer(ibChan chan inputBatch) (chan resultBatch, chan bool){
	out := make(chan resultBatch, bufSize)
	counts := make(chan bool)

	go func(){

		_, env, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{passfaultClassPath}))
		check(err)

		obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
		check(err)

		for ib := range ibChan{

			rb := resultBatch{}

			// range over inputBatch
			for _, input := range ib{
				freq := input.freq
				pass := input.pass

				str, err := env.NewObject("java/lang/String", []byte(pass))
				check(err)

				// call passwordAnalysis on password
				v, err := obj.CallMethod(env, "passwordAnalysis", "java/lang/String", str)
				check(err)

				// format result from JVM into byte array (probably not the most elegant way!)
				resultJVM, err := v.(*jnigi.ObjectRef).CallMethod(env, "getBytes", jnigi.Byte|jnigi.Array)
				resultString := string(resultJVM.([]byte))
				resultBytes := []byte(resultString)

				rb = append(rb, result{freq, resultBytes})

				// signal to counter
				counts <- true
			}

			out <- rb
		}

		close(out)
		close(counts)
	}()

	return out, counts
}

func counter(c *int, counts chan bool){
	for _ = range counts{
		*c++
	}
}

func reporter(c *int, total int, done bool){
	start := time.Now()

	lastCount := 0
	// report progress every second
	for !done{
		time.Sleep(1000 * time.Millisecond)

		since := time.Since(start)
		speed := *c - lastCount
		progress := float64(*c*100)/float64(total)

		fmt.Println("Speed: " + strconv.Itoa(speed) + " P/s; Progress: " + strconv.FormatFloat(progress, 'f', 2, 64) + " %; Processed passwords: " + strconv.Itoa(*c) + "; Total time: " + since.String())
		lastCount = *c
	}
}

func batchDecoder(rbChan chan resultBatch) chan trainedMaps{
	tmChan := make(chan trainedMaps, bufSize)

	go func(){
		for rb := range rbChan{
			tm := trainedMaps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

			for _, r := range rb{
				freq := r.freq
				result := r.result

				// parse JSON
				var cmFromJSON compositeJSON
				json.Unmarshal(result, &cmFromJSON)

				// update elementaryModel map
				for _, emFromJSON := range cmFromJSON.Models {
					if emFromMap, ok := tm.elementaries[emFromJSON.Name]; ok {	//elementary.Model already in tm, only update frequency
						emFromMap.UpdateTokenFreq(freq, emFromJSON.Token)
					}else{	// elementary.Model not in map, create new instance and insert into the map
						t := emFromJSON.Token
						f := freq

						tokenFreqs := make([]elementary.TokenFreq, 0)
						tokenFreqs = append(tokenFreqs, elementary.TokenFreq{t, f})

						newEM := elementary.Model{emFromJSON.Name, 0, tokenFreqs}
						tm.elementaries[emFromJSON.Name] = &newEM
					}
				}

				if cmFromMap, ok := tm.composites[cmFromJSON.CompositeModelName]; ok{	//composite.Model already in map, only update frequency
					cmFromMap.UpdateFreq(freq)
				}else{		// composite.Model not in map, create new instance and insert into the map
					compModelName := cmFromJSON.CompositeModelName

					elementaryModels := make([]string, 0)
					for _, emFromJSON := range cmFromJSON.Models {
						elementaryModels = append(elementaryModels, emFromJSON.Name)
					}

					// instantiate new Composite Model
					//cm := composite.Model{compModelName, complexity, freq, 0, elementaryModels}
					cm := composite.Model{compModelName, freq, 0, 0, elementaryModels}

					// add to map
					tm.composites[compModelName] = &cm
				}
			}

			tmChan <- tm
		}

		close(tmChan)
	}()

	return tmChan
}

func mapsMerger(tmChan chan trainedMaps) trainedMaps{
	finalMaps := trainedMaps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

	for tm := range tmChan{

		// process tm.elementaries
		for k, emFrom := range tm.elementaries {
			if emTo, ok := finalMaps.elementaries[k]; ok{ //emFrom already in finalMaps.elementaries

				//verify every TokenFreq in emFrom
				for _, tf := range emFrom.TokenFreqs{
					token := tf.Token
					freq := tf.Freq

					// is t already in emTo.TokenFreqs??
					index := 0
					b := false
					for i, tfTo := range emTo.TokenFreqs{
						if tfTo.Token == token{
							index = i
							b = true
							break
						}
					}

					if b{		// yes, t is in emTo.TokensNfreqs
						emTo.TokenFreqs[index].Freq += freq
					}else{		//no, t is not in emTo.TokensNfreqs
						emTo.TokenFreqs = append(emTo.TokenFreqs, elementary.TokenFreq{token, freq})
					}

					emTo.Sort()
				}
			}else{			//emFrom not in finalMaps.elementaries, create new entry
				emFrom.Sort()
				finalMaps.elementaries[k] = emFrom
			}
		}

		// process composite.Models
		for k, cmFrom := range tm.composites {
			if cmTo, ok := finalMaps.composites[k]; ok{ //cm already in finalMaps.composites
				cmTo.UpdateFreq(cmFrom.Freq)
			}else{		//cm not in finalMaps.composites, create new entry

				//insert cm
				finalMaps.composites[k] = cmFrom
			}
		}
	}

	for _, em := range finalMaps.elementaries{
		em.UpdateEntropy()
	}

	sum := 0
	for _, cm := range finalMaps.composites{
		sum += cm.Freq
	}

	for _, cm := range finalMaps.composites{
		cm.UpdateProb(sum)
	}

	return finalMaps
}

func saveMaps(finalMaps trainedMaps, list string){
	fmt.Println("Saving maps...")
	emArray := make([]elementary.Model, len(finalMaps.elementaries))
	i := 0
	for _, em := range finalMaps.elementaries{
		emArray[i] = *em
		i++
	}
	jsonData, err := json.Marshal(emArray)

	emFile, err := os.Create("maps/" + list + "Elementaries.json")
	check(err)
	defer emFile.Close()

	emFile.Write(jsonData)
	emFile.Sync()

	cmArray := make([]composite.Model, len(finalMaps.composites))
	i = 0
	for _, cm := range finalMaps.composites{

		// ignore improbable cms
		if cm.Prob > 0.00001{
			cmArray[i] = *cm
			i++
		}
	}

	jsonData, err = json.Marshal(cmArray)

	cmFile, err := os.Create("maps/" + list + "Composites.json")
	check(err)
	defer cmFile.Close()

	cmFile.Write(jsonData)
	cmFile.Sync()
}

func Train(list string){
	total, inputBatches := generator(list, 100)
	resultBatches, counts := batchAnalyzer(inputBatches)
	c := 0
	go counter(&c, counts)
	go reporter(&c, total, false)
	trainedMaps := batchDecoder(resultBatches)
	finalMaps := mapsMerger(trainedMaps)
	saveMaps(finalMaps, list)
	os.Exit(0)
}

