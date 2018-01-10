package main

import (
	"os"
	"compress/gzip"
	"io/ioutil"
	"github.com/bernardoaraujor/corinda/composite"
	"encoding/json"
	"runtime"
	"fmt"
	"github.com/bernardoaraujor/corinda/elementary"
)

func main() {
	checkComplexity("rockyou")
	checkComplexity("linkedin")
	checkComplexity("antipublic")

	checkComplexity("rockyou_1M")
	checkComplexity("linkedin_1M")
	checkComplexity("antipublic_1M")
}

func checkComplexity(list string){

	// load composites into array
	f, err := os.Open("maps/" + list + "Composites.json.gz")
	check(err)
	defer f.Close()
	gr, err := gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cmJSON, err := ioutil.ReadAll(gr)
	check(err)
	var composites1 []*composite.Model
	err = json.Unmarshal(cmJSON, &composites1)
	check(err)

	// load elementaries into array
	f, err = os.Open("maps/" + list + "Elementaries.json.gz")
	check(err)
	defer f.Close()
	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	emJSON, err := ioutil.ReadAll(gr)
	check(err)
	var elementaries1 []*elementary.Model
	err = json.Unmarshal(emJSON, &elementaries1)
	check(err)

	// load elementaries into map
	emMap := make(map[string]*elementary.Model)
	for _, em := range elementaries1{
		emMap[em.Name] = em
	}

	// calculate complexities
	complexities := make([]int, 0)
	for _, cm := range composites1{
		complexity := 1
		for _, em := range cm.Models{
			complexity *= len(emMap[em].TokenFreqs)
		}

		complexities = append(complexities, complexity)
	}

	output, err := os.Create("statistics/" + list + "_comp.txt")
	check(err)
	defer output.Close()

	for _, complexity := range complexities{
		fmt.Fprintln(output, complexity)
	}
}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

