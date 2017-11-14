package main

import (
	"github.com/bernardoaraujor/corinda"
	"os"
	"strconv"
	"time"
	"fmt"
)

func main() {
	// path for input file is first arg
	inputPath := os.Args[1]

	// start JVM
	jvm := corinda.StartJVM()

	// number of PasswordAnalysis go routines is second arg
	nRoutines, _ := strconv.Atoi(os.Args[2])

	// start counter
	count := 0
	countChan := make(chan bool, nRoutines)
	go corinda.Counter(&count, countChan)

	// channels for processing passwords
	fpChan := make(chan corinda.FreqNpass, nRoutines)
	frChan := make(chan corinda.FreqNresult, nRoutines)

	// launch csv reader
	go corinda.CsvRead(inputPath, fpChan)

	// start go routines for processing passwords
	for i := 0; i < nRoutines; i++ {
		go corinda.PasswordAnalysis(jvm, fpChan, frChan, countChan)
	}

	// start go routine for decoding json and saving data to maps
	go corinda.DecodeJSON(frChan)

	// report progress every second
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Processed passwords: " + strconv.Itoa(count))
	}
}



