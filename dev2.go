package main

import (
	"github.com/timob/jnigi"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type freqNpass struct{
	freq int
	pass string
}

func main() {

	//TODO: how to get slices from .csv.gz??
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

	//TODO: buffer this channel... which size?
	fpChan := make(chan freqNpass)

	go csvRead(gr, done, fpChan)

	//TODO: many processes... how to parametrize in function of n?
	go processPass(fpChan)

	for !<-done{

	}
}

func csvRead(gr *gzip.Reader, done chan bool, fpChan chan freqNpass) {
	r := csv.NewReader(gr)

	for records, err := r.Read(); records != nil; records, err = r.Read(){
		if err != nil {
			log.Fatal(err)
		}

		freq, err := strconv.Atoi(records[0])
		pass := records[1]
		if err != nil {
			log.Fatal(err)
		}

		fp := freqNpass{freq, pass}
		fpChan <- fp
	}

	done <- true
}

func processPass(fpChan chan freqNpass){
	//start JVM
	_, env, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{"-Xcheck:jni", "-Djava.class.path=/home/bernardo/go/src/github.com/bernardoaraujor/corinda/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"}))
	if err != nil {
		log.Fatal(err)
	}

	//create TextAnalysis JVM object
	obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
	if err != nil {
		log.Fatal(err)
	}

	//iterate over channel
	//TODO: close channel?
	for {
		//read channel
		fp, ok := <-fpChan
		if !ok {
			break
		}

		//create JVM string with password
		str, err := env.NewObject("java/lang/String", []byte(fp.pass))

		//call passwordAnalysis on password
		v, err := obj.CallMethod(env, "passwordAnalysis", "java/lang/String", str)
		if err != nil {
			log.Fatal(err)
		}

		//format result from JVM into Go string
		resultJVM, err := v.(*jnigi.ObjectRef).CallMethod(env, "getBytes", jnigi.Byte|jnigi.Array)
		resultGo := string(resultJVM.([]byte))

		fmt.Println(fp.pass + " " + resultGo)
	}
}

