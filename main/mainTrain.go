package main

import (
	"github.com/bernardoaraujor/corinda/train"
	"os"
)

func main() {
	list := os.Args[1]

	train.Train(list)
}