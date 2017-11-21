package main

import (
	"github.com/bernardoaraujor/corinda/train"
	//"os"
	//"strconv"
	//"testing"
)

func main() {
	//input := os.Args[1]
	//nRoutines, _ := strconv.Atoi(os.Args[2])

	input := "test"
	nRoutines := 1
	train.Train(input, nRoutines)
}