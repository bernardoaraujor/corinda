package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"os"
)

func main() {
	input := os.Args[1]

	train.Train(input)
}