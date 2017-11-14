package corinda

import (
	"github.com/timob/jnigi"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"encoding/json"
	//"fmt"
	"fmt"
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

func StartJVM() (*jnigi.JVM){
	jvm, _, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{"-Xcheck:jni", "-Djava.class.path=/home/bernardo/go/src/github.com/bernardoaraujor/corinda/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"}))
	if err != nil {
		log.Fatal(err)
	}

	return jvm
}

func CsvRead(f *os.File, done chan bool, fpChan chan FreqNpass) {
	r := csv.NewReader(f)

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
	close(fpChan)

	done <- true
}

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

		/*
		for _, b := range compModel.ElementaryModels{
			fmt.Println(b)
		}
		*/
	}
}

func ProcessPass(jvm *jnigi.JVM, fpChan chan FreqNpass, frChan chan FreqNresult, countChan chan bool){
	env := jvm.AttachCurrentThread()

	//create TextAnalysis JVM object
	obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
	if err != nil {
		log.Fatal(err)
	}


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

		frChan <- FreqNresult{fp.freq, resultBytes}
		countChan <- true
	}
	close(frChan)
}

func Counter(c *int, countChan chan bool){
	for _ = range countChan{
		*c++
	}
}