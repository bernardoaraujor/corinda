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

func main(){
	tm := train.Maps{make(map[string]*elementary.Model), make(map[string]*composite.Model)}

	var tm2 = new(train.Maps)
	err := load("maps/testTrainedMaps.gob", &tm2)
	check(err)
	tm.Merge(tm2)

	em := tm.ElementaryMap["Exact Match:10k-worst-passwords"]
	tf := em.TokensNfreqs

	f, err := os.Create("export.txt")
	check(err)
	defer f.Close()

	for _, v := range tf{
		fmt.Println(v.Freq)
		fmt.Fprintln(f, v.Freq)
	}
}

/*
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