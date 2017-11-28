package main

import (
	"github.com/bernardoaraujor/corinda/crack"
	"fmt"
	"runtime"
	"os"
	"encoding/gob"
)

func main(){
	crack := crack.NewCrack("test", crack.SHA1)

	crack.Crack()
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