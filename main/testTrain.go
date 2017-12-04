package main

import (
	"github.com/bernardoaraujor/corinda/train"
	//"os"
	//"strconv"
	//"testing"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	nRoutines, _ := strconv.Atoi(os.Args[2])

	//input := "rockyou"
	//nRoutines := 10
	train.Train(input, nRoutines)
}