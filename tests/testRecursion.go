package main

import (
	"github.com/bernardoaraujor/corinda/composite"
	"os"
	"compress/gzip"
	"io/ioutil"
	"encoding/json"
	"runtime"
	"fmt"
	"github.com/bernardoaraujor/corinda/elementary"
)

func main() {

	list := "rockyou"

	f, err := os.Open("maps/"+ list +"Elementaries.json.gz")
	check(err)

	gr, err := gzip.NewReader(f)
	check(err)

	raw, err := ioutil.ReadAll(gr)
	check(err)
	var elementaries []*elementary.Model
	err = json.Unmarshal(raw, &elementaries)
	check(err)

	emMap := make(map[string]*elementary.Model)

	for _, em := range elementaries{
		emMap[em.Name] = em
	}

	f, err = os.Open("maps/"+ list +"Composites.json.gz")
	check(err)

	gr, err = gzip.NewReader(f)
	check(err)

	raw, err = ioutil.ReadAll(gr)
	check(err)
	var composites []*composite.Model
	err = json.Unmarshal(raw, &composites)

	cmMap := make(map[string]*composite.Model)
	for _, cm := range composites{
		cmMap[cm.Name] = cm
	}

	cm := cmMap["|5 Random Character(s):Numbers|Exact Match:JohnTheRipper.txt|"]
	tokenLists := make([][]string, 0)
	for _, elementaryName := range cm.Models{
		elementary := emMap[elementaryName]

		tokenLists = append(tokenLists, elementary.SortedTokens())
	}
	for guess := range cm.Guess(tokenLists){
		fmt.Println(guess)
	}

}


func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}