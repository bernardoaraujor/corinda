package main

import (
	"github.com/bernardoaraujor/corinda"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"fmt"
)

func main() {
	inputPath := os.Args[1]
	in, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	outputPath := strings.Replace(inputPath, ".csv", "_out.csv", -1)
	out, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	//start JVM
	jvm := corinda.StartJVM()

	nRoutines, _ := strconv.Atoi(os.Args[2])

	count := 0
	countChan := make(chan bool, nRoutines)
	go corinda.Counter(&count, countChan)

	done := make(chan bool)
	fpChan := make(chan corinda.FreqNpass, nRoutines)
	frChan := make(chan corinda.FreqNresult, nRoutines)

	go corinda.CsvRead(in, done, fpChan)
		for i := 0; i < nRoutines; i++ {
		go corinda.ProcessPass(jvm, fpChan, frChan, countChan)
	}

	//go corinda.CsvWrite(out, frChan)
	go corinda.DecodeJSON(frChan)

	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Processed passwords: " + strconv.Itoa(count))
	}
}



