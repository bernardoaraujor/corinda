package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"fmt"
	"os"
	"encoding/gob"
	"runtime"
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/bernardoaraujor/corinda/composite"
)

func main() {
	// empty tm
	tm := train.Maps{make(map[string]*elementary.Model), make(map[string]*composite.Model), 0}

	em := tm.ElementaryMap[""]
	fmt.Println(em)

	//train.Train("test2", 1)
	var tm2 = new(train.Maps)
	err := Load("maps/testTrainedMaps.gob", &tm2)
	Check(err)

	tm.Merge(tm2)

	for _, v := range tm.ElementaryMap {
		if v.Name == ""{
			fmt.Println(0)
		}
	}

}

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}