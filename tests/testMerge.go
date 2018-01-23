package main

import (
	//"github.com/bernardoaraujor/corinda/train"
	"fmt"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
	"io/ioutil"
	"encoding/json"

	"runtime"
	"os"
	"compress/gzip"
	"github.com/bernardoaraujor/corinda/train"
)

func main() {

	f, err := os.Open("maps/rockyou1Elementaries.json.gz")
	check(err)
	defer f.Close()
	gr, err := gzip.NewReader(f)
	check(err)
	defer gr.Close()

	em1, err := ioutil.ReadAll(gr)
	check(err)
	var elementaries []*elementary.Model
	err = json.Unmarshal(em1, &elementaries)
	check(err)

	emMap1 := make(map[string]*elementary.Model)
	for _, em := range elementaries{
		emMap1[em.Name] = em
	}

	f, err = os.Open("maps/rockyou1Composites.json.gz")
	check(err)
	defer f.Close()
	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cm1, err := ioutil.ReadAll(gr)
	check(err)
	var composites []*composite.Model
	err = json.Unmarshal(cm1, &composites)
	check(err)

	cmMap1 := make(map[string]*composite.Model)
	for _, cm := range composites{
		cmMap1[cm.Name] = cm
	}

	//------------------------------

	f, err = os.Open("maps/rockyou2Elementaries.json.gz")
	check(err)
	defer f.Close()
	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	em2, err := ioutil.ReadAll(gr)
	check(err)
	var elementaries2 []*elementary.Model
	err = json.Unmarshal(em2, &elementaries2)
	check(err)

	emMap2 := make(map[string]*elementary.Model)
	for _, em := range elementaries{
		emMap2[em.Name] = em
	}

	f, err = os.Open("maps/rockyou2Composites.json.gz")
	check(err)
	defer f.Close()
	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cm2, err := ioutil.ReadAll(gr)
	check(err)
	var composites2 []*composite.Model
	err = json.Unmarshal(cm2, &composites2)
	check(err)

	cmMap2 := make(map[string]*composite.Model)
	for _, cm := range composites{
		cmMap2[cm.Name] = cm
	}

	fmt.Println(len(cmMap2))
	train.Merge(emMap1, emMap2, cmMap1, cmMap2)
	fmt.Println(len(cmMap2))

	//------------------------------------------------------

	emArray := make([]elementary.Model, len(emMap2))
	i := 0
	for _, em := range emMap2{
		emArray[i] = *em
		i++
	}
	jsonData, err := json.Marshal(emArray)

	emFile, err := os.Create("maps/rockyouMergedElementaries.json")
	check(err)
	defer emFile.Close()

	emFile.Write(jsonData)
	emFile.Sync()

	cmArray := make([]composite.Model, len(cmMap2))
	i = 0
	for _, cm := range cmMap2{
		cmArray[i] = *cm
		i++
	}

	jsonData, err = json.Marshal(cmArray)

	cmFile, err := os.Create("maps/rockyouMergedComposites.json")
	check(err)
	defer cmFile.Close()

	cmFile.Write(jsonData)
	cmFile.Sync()
}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
