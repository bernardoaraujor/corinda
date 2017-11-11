package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

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
	//c := make(chan string)
	go readLines(gr, done)

	for !<-done{

	}

}
func readLines(gr *gzip.Reader, done chan bool) {
	r := csv.NewReader(gr)

	records, err := r.Read();

	for records != nil{
		if err != nil {
			log.Fatal(err)
		}

		freq := records[0]
		pass := records[1]
		fmt.Println(freq + " " + pass)
		records, err = r.Read()
	}

	done <- true
}
