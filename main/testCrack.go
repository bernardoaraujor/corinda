package main

import (
	"github.com/bernardoaraujor/corinda/crack"
	"fmt"
)

func main(){
	crack := crack.NewCrack("test", crack.SHA1)

	batches := crack.Crack()

	batch := <- batches
	fmt.Println(batch)
}