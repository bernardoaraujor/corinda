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
	tm := train.TrainedMaps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

	var tm2 = new(train.TrainedMaps)
	err := load("maps/testTrainedMaps.gob", &tm2)
	check(err)
	tm.Merge(tm2)

	cm := tm.CompositeModelsMap["|Exact Match:frPopular.txt|Exact Match:500-worst-passwords.txt|"]

	c := cm.Guess()

	i := 0
	for s := range c{
		i++
		fmt.Println(s)
	}
	fmt.Println(i)
}

