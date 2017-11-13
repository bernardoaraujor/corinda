package main

import (
	"github.com/timob/jnigi"
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	//"fmt"
	"fmt"
)

type freqNpass struct{
	freq int
	pass string
}

type freqNresult struct{
	freq int
	result string
}

func startJVM() (*jnigi.JVM){
	jvm, _, err := jnigi.CreateJVM(jnigi.NewJVMInitArgs(false, true, jnigi.DEFAULT_VERSION, []string{"-Xcheck:jni", "-Djava.class.path=/home/bernardo/go/src/github.com/bernardoaraujor/corinda/passfault_corinda/out/artifacts/passfault_corinda_jar/passfault_corinda.jar"}))
	if err != nil {
		log.Fatal(err)
	}

	return jvm
}

func csvRead(f *os.File, done chan bool, fpChan chan freqNpass) {
	r := csv.NewReader(f)

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
	close(fpChan)

	done <- true
}

func csvWrite(f *os.File, frChan chan freqNresult){
	writer := csv.NewWriter(f)
	defer writer.Flush()

	for fr := range frChan {
		s := []string{strconv.Itoa(fr.freq), fr.result}
		err := writer.Write(s)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func processPass(jvm *jnigi.JVM, fpChan chan freqNpass, frChan chan freqNresult, countChan chan bool){
	env := jvm.AttachCurrentThread()

	//create TextAnalysis JVM object
	obj, err := env.NewObject("org/owasp/passfault/TextAnalysis")
	if err != nil {
		log.Fatal(err)
	}


	for fp := range fpChan {
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
		frChan <- freqNresult{fp.freq, resultGo}
		countChan <- true
	}
	close(frChan)
}

func counter(c *int, countChan chan bool){
	for _ = range countChan{
		*c++
	}
}

func main() {
	inputPath := os.Args[1]
	in, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	outputPath := strings.Replace(inputPath, ".csv", "_out.csv", -1)
	out, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	//start JVM
	jvm := startJVM()

	done := make(chan bool)

	nRoutines, _ := strconv.Atoi(os.Args[2])

	count := 0
	countChan := make(chan bool, nRoutines)

	go counter(&count, countChan)
	fpChan := make(chan freqNpass, nRoutines)
	frChan := make(chan freqNresult, nRoutines)

	go csvRead(in, done, fpChan)
		for i := 0; i < nRoutines; i++ {
		go processPass(jvm, fpChan, frChan, countChan)
	}

	go csvWrite(out, frChan)

	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Processed passwords: " + strconv.Itoa(count))
	}
}



