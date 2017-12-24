package main

import (
	"io/ioutil"
	"github.com/bernardoaraujor/corinda/elementary"
	"encoding/json"
	"github.com/bernardoaraujor/corinda/composite"
	"fmt"
	"runtime"
	"os"
	"encoding/gob"
	"compress/gzip"
)

func main() {
	list := "rockyou"

	f, err := os.Open("maps/"+ list +"Elementaries.json.gz")
	check(err)

	gr, err := gzip.NewReader(f)
	check(err)

	raw, err := ioutil.ReadAll(gr)
	check(err)
	var elementaries []*elementary.Model
	err = json.Unmarshal(raw, &elementaries)
	check(err)

	elementariesMap := make(map[string]*elementary.Model)

	for _, em := range elementaries{
		sortedTokens := em.SortedTokens()
		fmt.Println(em.TokenFreqs)
		fmt.Println(sortedTokens)
		elementariesMap[em.Name] = em
	}

	raw, err = ioutil.ReadFile("maps/"+ list +"Composites.json")
	check(err)
	var composites []*composite.Model
	err = json.Unmarshal(raw, &composites)

	compositesMap := make(map[string]*composite.Model)

	for _, cm := range composites{
		compositesMap[cm.Name] = cm
	}

}


