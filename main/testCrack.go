package main

import (
	"github.com/bernardoaraujor/corinda/crack"
)

func main(){
	crack := crack.NewCrack("test", crack.SHA1)

	crack.Crack()
}