package main

import (
	"fmt"
	"os"
	"encoding/gob"
	"runtime"
	"github.com/bernardoaraujor/corinda/train"
)

type User struct {
	Name, Pass string
}

func main(){
	var tm = new(train.TrainedMaps)
	err := Load("maps/xaaTrainedMaps.gob", tm)
	Check(err)

	var tm2 = new(train.TrainedMaps)
	err = Load("maps/xabTrainedMaps.gob", tm2)
	Check(err)

	fmt.Println(tm.CompositeModelsMap["|Backwards: Exact Match:nlPopular|3 Keyboard Repeated Character(s):English Keyboard|"])
	fmt.Println(tm2.CompositeModelsMap["|Backwards: Exact Match:nlPopular|3 Keyboard Repeated Character(s):English Keyboard|"])
	tm.Merge(tm2)
	fmt.Println(tm.CompositeModelsMap["|Backwards: Exact Match:nlPopular|3 Keyboard Repeated Character(s):English Keyboard|"])
}

// Encode via Gob to file
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
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