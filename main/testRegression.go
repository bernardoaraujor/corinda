package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
	"runtime"
	"fmt"
	"os"
	"encoding/gob"
	"strconv"
)

func main() {
	tm := train.TrainedMaps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

	var tm2 = new(train.TrainedMaps)
	err := load("maps/testTrainedMaps.gob", &tm2)
	check(err)
	tm.Merge(tm2)

	for _, em := range tm.ElementaryModelsMap{
		em.UpdateEntropy()
		fmt.Println(strconv.FormatFloat(em.Entropy, 'f', -1, 64) + "," + strconv.Itoa(em.Complexity) + "," + strconv.Itoa(len(em.TokensNfreqs)))
	}
	a := 1
	fmt.Println(a)
}


// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

// Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}