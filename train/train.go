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

func batchAnalyzer(c *int, ibChan chan inputBatch) chan resultBatch{
	out := make(chan resultBatch, bufSize)

	go func(){

		jvm, _, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{passfaultClassPath}))
		check(err)

		for ib := range ibChan{

			// attach this routine to JVM
			env := jvm.AttachCurrentThread()

			obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
			check(err)

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

				// increment counter
				*c++

				//env.DeleteGlobalRef(str)
				env.DeleteLocalRef(str)
			}

			out <- rb

			jvm.DetachCurrentThread()
		}

		close(out)
	}()

	return out
}

func reporter(m *int, c *int, total int, done bool){
	start := time.Now()

	lastCount := 0
	// report progress every second
	for !done{
		time.Sleep(1000 * time.Millisecond)

		since := time.Since(start)
		speed := *c - lastCount
		progress := float64(*c*100)/float64(total)

		fmt.Println("Maps merged: " + strconv.Itoa(*m) + "; Speed: " + strconv.Itoa(speed) + " P/s; Progress: " + strconv.FormatFloat(progress, 'f', 2, 64) + " %; Processed passwords: " + strconv.Itoa(*c) + "; Total time: " + since.String())
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

						tokenFreqs := make(map[string]int)
						tokenFreqs[t] = f

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

func mapsMerger(m *int, tmChan chan trainedMaps) trainedMaps{
	finalMaps := trainedMaps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

	for tm := range tmChan{

		// process tm.elementaries
		for k, emFrom := range tm.elementaries {
			if emTo, ok := finalMaps.elementaries[k]; ok{ //emFrom already in finalMaps.elementaries

				//verify every TokenFreq in emFrom
				for token, freq := range emFrom.TokenFreqs{

					// is t already in emTo.TokenFreqs??
					if _, ok := emTo.TokenFreqs[token]; ok{
						f := emTo.TokenFreqs[token]
						emTo.TokenFreqs[token] = f + freq
					}else{
						emTo.TokenFreqs[token] = freq
					}
				}

			}else{			//emFrom not in finalMaps.elementaries, create new entry

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

		*m++
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
		cm.UpdateEntropy(finalMaps.elementaries)
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
	total, inputBatches := generator(list, 1000000)
	c := 0
	m := 0
	resultBatches := batchAnalyzer(&c, inputBatches)
	go reporter(&m, &c, total, false)
	trainedMaps := batchDecoder(resultBatches)
	finalMaps := mapsMerger(&m, trainedMaps)
	saveMaps(finalMaps, list)
	os.Exit(0)
}

