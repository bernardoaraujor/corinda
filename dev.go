/*
package main


import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type csvLine struct{
	freq int
	pass string
}

func main() {
	f, err := os.Open("lists/rockyou.csv.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()

	done := make(chan bool)
	csvLineChan := make(chan csvLine)
	go csvRead(gr, csvLineChan, done)

	for !<-done{
		csvLine := <-csvLineChan
		fmt.Println(strconv.Itoa(csvLine.freq) + " " + csvLine.pass)
	}

}
func csvRead(gr *gzip.Reader, csvLineChan chan csvLine, done chan bool) {
	r := csv.NewReader(gr)

	for records, err := r.Read(); records != nil; records, err = r.Read(){
		if err != nil {
			log.Fatal(err)
		}

		freq, err := strconv.Atoi(records[0])
		pass := records[1]
		if err != nil {
			log.Fatal(err)
		}

		line := csvLine{freq, pass}
		//fmt.Println(strconv.Itoa(line.freq) + " " + line.pass)
		csvLineChan <- line
	}

	done <- true
}
 */