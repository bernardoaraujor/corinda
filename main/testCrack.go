package main

import (
	"github.com/bernardoaraujor/corinda/crack"
	"os"
)

func main(){
	trained := os.Args[1]
	target := os.Args[2]
	alg := os.Args[3]

	algs := ""
	if alg == "sha1"{
		algs = crack.SHA1
	}else if alg == "sha256"{
		algs = crack.SHA256
	}

	crack := crack.NewCrack(trained, target, algs, 1.0)

	crack.Crack()
}