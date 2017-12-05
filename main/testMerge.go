package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"fmt"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
	"io/ioutil"
	"encoding/json"

	"runtime"
	"os"
)

func main() {
	raw, err := ioutil.ReadFile("maps/rockyou10kElementaries.json")
	check(err)
	var elementaries []*elementary.Model
	err = json.Unmarshal(raw, &elementaries)
	check(err)

	emMap1 := make(map[string]*elementary.Model)
	for _, em := range elementaries{
		emMap1[em.Name] = em
	}

	raw, err = ioutil.ReadFile("maps/rockyou10kComposites.json")
	check(err)
	var composites []*composite.Model
	err = json.Unmarshal(raw, &composites)
	check(err)

	cmMap1 := make(map[string]*composite.Model)
	for _, cm := range composites{
		cmMap1[cm.Name] = cm
	}

	//-----------
	raw, err = ioutil.ReadFile("maps/rockyou10k-bElementaries.json")
	check(err)
	var elementaries2 []*elementary.Model
	err = json.Unmarshal(raw, &elementaries2)
	check(err)

	emMap2 := make(map[string]*elementary.Model)
	for _, em := range elementaries2{
		emMap2[em.Name] = em
	}

	raw, err = ioutil.ReadFile("maps/rockyou10k-bComposites.json")
	check(err)
	var composites2 []*composite.Model
	err = json.Unmarshal(raw, &composites2)
	check(err)

	cmMap2 := make(map[string]*composite.Model)
	for _, cm := range composites2{
		cmMap2[cm.Name] = cm
	}

	//-------------------------
	train.Merge(emMap1, emMap2, cmMap1, cmMap2)

}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
