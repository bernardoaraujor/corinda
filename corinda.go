package corinda

import (
	"github.com/timob/jnigi"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"encoding/json"
	//"fmt"
)

type FreqNpass struct{
	freq int
	pass string
}

type FreqNresult struct{
	freq int
	result []byte
}


//these types are later converted into gob files
//they represent conrinda's "training"
type ElementaryModel struct{
	modelName string
	complexity int
	tokenFreqMap map[string]int
}

func (em ElementaryModel) updateTokenFreqMap(freq int, token string){
	if freqFromMap, ok := em.tokenFreqMap[token]; ok {	//token already in map, only update frequency
		em.tokenFreqMap[token] = freqFromMap + freq
	}else{	//token not in map, insert new freq into map
		em.tokenFreqMap[token] = freq
	}
}

type CompositeModel struct{
	compModelName string
	complexity int
	freq int
	elementaryModels []*ElementaryModel
}

func (cm CompositeModel) updateFreq(freq int){
	cm.freq = cm.freq + freq
}

//the types ElModelJSON and CompModelJSON are used only for parsing JSON into ElementaryModel and CompositeModel
//this is done by function DecodeJSON
type ElModelJSON struct{
	ModelName  string `json:"modelName"`
	Complexity int    `json:"complexity"`
	ModelIndex int    `json:"modelIndex"`
	Token      string `json:"token"`
}

type CompModelJSON struct {
	Complexity       int `json:"complexity"`
	ElementaryModels []ElModelJSON `json:"elementaryModels"`
	CompositeModelName string `json:"compositeModelName"`
}

//starts JVM that will be used to call Passfault's passwordAnalysis
func StartJVM() (*jnigi.JVM){
	jvm, _, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{"-Djava.class.path=/home/bernardo/go/src/github.com/bernardoaraujor/corinda/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"}))
	if err != nil {
		log.Fatal(err)
	}

	return jvm
}

//reads lines from csv file and sends them to a buffered channel
//many go routines of ProcessPass will read from this channel
func CsvRead(inputPath string, fpChan chan FreqNpass) {
	in, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	r := csv.NewReader(in)

	for records, err := r.Read(); records != nil; records, err = r.Read(){
		if err != nil {
			log.Fatal(err)
		}

		freq, err := strconv.Atoi(records[0])
		pass := records[1]
		if err != nil {
			log.Fatal(err)
		}

		fp := FreqNpass{freq, pass}
		fpChan <- fp
	}

	//no more lines, close channel
	close(fpChan)
}

//parses the JSON strings returned from Passfault
//data is stored in maps of Composite and Elementary Models
func DecodeJSON(frChan chan FreqNresult){
	compositeModelMap := make(map[string]CompositeModel)
	elementaryModelMap := make(map[string]ElementaryModel)
	for fr := range frChan{
		freq := fr.freq
		result := fr.result

		//parse JSON
		var cmFromJSON CompModelJSON
		json.Unmarshal(result, &cmFromJSON)

		//update elementaryModel map
		for _, emFromJSON := range cmFromJSON.ElementaryModels{
			if emFromMap, ok := elementaryModelMap[emFromJSON.ModelName]; ok {	//ElementaryModel already in map, only update frequency
				emFromMap.updateTokenFreqMap(freq, emFromJSON.Token)
			}else{	//ElementaryModel not in map, create new instance and insert into the map
				tokenFreqMap := make(map[string]int)
				tokenFreqMap[emFromJSON.Token] = freq
				newEM := ElementaryModel{emFromJSON.ModelName, emFromJSON.Complexity, tokenFreqMap}
				elementaryModelMap[emFromJSON.ModelName] = newEM
			}
		}

		if cmFromMap, ok := compositeModelMap[cmFromJSON.CompositeModelName]; ok{	//CompositeModel already in map, only update frequency
			cmFromMap.updateFreq(freq)
		}else{		//CompositeModel not in map, create new instance and insert into the map
			compModelName := cmFromJSON.CompositeModelName
			complexity := cmFromJSON.Complexity

			//populate array of pointers
			var elementaryModels []*ElementaryModel
			for _, emFromJSON := range cmFromJSON.ElementaryModels{
				emName := emFromJSON.ModelName
				em := elementaryModelMap[emName]
				elementaryModels = append(elementaryModels, &em)
			}

			//instantiate new Composite Model
			cm := CompositeModel{compModelName, complexity, freq, elementaryModels}

			//add to map
			compositeModelMap[compModelName] = cm
		}
	}
}

//uses Passfault's passwordAnalysis method to process passwords
func PasswordAnalysis(jvm *jnigi.JVM, fpChan chan FreqNpass, frChan chan FreqNresult, countChan chan bool){
	//attach this routine to JVM
	env := jvm.AttachCurrentThread()

	//create TextAnalysis JVM object
	obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
	if err != nil {
		log.Fatal(err)
	}

	//loop over inputs from csv file
	for fp := range fpChan {
		//create JVM string with password
		str, err := env.NewObject("java/lang/String", []byte(fp.pass))

		//call passwordAnalysis on password
		v, err := obj.CallMethod(env, "passwordAnalysis", "java/lang/String", str)
		if err != nil {
			log.Fatal(err)
		}

		//format result from JVM into byte array (probably not the most elegant way!)
		resultJVM, err := v.(*jnigi.ObjectRef).CallMethod(env, "getBytes", jnigi.Byte|jnigi.Array)
		resultString := string(resultJVM.([]byte))
		resultBytes := []byte(resultString)

		//send result to JSON decoder
		frChan <- FreqNresult{fp.freq, resultBytes}

		//signal to counter
		countChan <- true
	}

	//no more inputs, buffer is empty, close channel
	close(frChan)
}

//counts how many passwords have already been analyzed
func Counter(c *int, countChan chan bool){
	for _ = range countChan{
		*c++
	}
}