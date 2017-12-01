package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
	"fmt"
	"runtime"
	"os"
	"encoding/gob"
)

func main() {

	// -----------------------------------------------------------------------------------------------------------------
	// empty tm
	tm := train.Maps{make(map[string]*elementary.Model), make(map[string]*composite.Model), 0}

	var tm2 = new(train.Maps)
	err := load("maps/testTrainedMaps.gob", &tm2)
	check(err)
	tm.Merge(tm2)

	cm := tm.CompositeMap["|Exact Match:frPopular.txt|Exact Match:500-worst-passwords.txt|"]

	c := cm.Guess()

	for s := range c{
		fmt.Println(s)
	}
}

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

