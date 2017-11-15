package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	nRoutines, _ := strconv.Atoi(os.Args[2])
	train.Train(input, nRoutines)
}