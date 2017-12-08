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
)

func main() {
	em1, err := ioutil.ReadFile("maps/rockyou1Elementaries.json")
	check(err)
	var elementaries []*elementary.Model
	err = json.Unmarshal(em1, &elementaries)
	check(err)

	emMap1 := make(map[string]*elementary.Model)
	for _, em := range elementaries{
		emMap1[em.Name] = em
	}

	cm1, err := ioutil.ReadFile("maps/rockyou1Composites.json")
	check(err)
	var composites []*composite.Model
	err = json.Unmarshal(cm1, &composites)
	check(err)

	cmMap1 := make(map[string]*composite.Model)
	for _, cm := range composites{
		cmMap1[cm.Name] = cm
	}

	fmt.Println(0)
}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
