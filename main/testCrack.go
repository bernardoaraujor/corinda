package main

import (
	//"github.com/bernardoaraujor/corinda/train"
	"os"
	"fmt"
	"runtime"
	"log"
	"io/ioutil"
	"github.com/bernardoaraujor/corinda/train"
	"encoding/gob"
	"github.com/bernardoaraujor/corinda/crack"
)

func mergeMaps() train.TrainedMaps{
	// empty tm
	tm := train.TrainedMaps{make(map[string]*train.ElementaryModel), make(map[string]*train.CompositeModel)}

	files, err := ioutil.ReadDir("maps")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//train.Train("test2", 1)
		var tm2 = new(train.TrainedMaps)
		err := load("maps/" + f.Name(), &tm2)
		check(err)

		tm.Merge(tm2)
	}

	return tm
}

func main(){
	tm := mergeMaps()
	cm := tm.CompositeModelsMap["|Exact Match:JohnTheRipper|"]

	sum := 0
	for _, v := range tm.CompositeModelsMap{
		sum += v.Freq
	}

	fmt.Println(sum)
	fmt.Println(cm.Freq)
	fmt.Println(int(tm.RelativeFreq(cm)*crack.MinBufferSize))
}