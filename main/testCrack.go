package main

import (
	"github.com/bernardoaraujor/corinda/crack"
)

func main(){
	crack := crack.NewCrack("rockyou10k", crack.SHA1)

	crack.Crack()
}