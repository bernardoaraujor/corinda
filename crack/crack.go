package crack

import (
	"runtime"
	"fmt"
	"os"
	//"github.com/bernardoaraujor/corinda/train"
	"github.com/bernardoaraujor/corinda/train"
	"encoding/gob"
)

const rockyou = 0
const linkedin = 1
const antipublic = 3

const MinBufferSize = 1000000

type Crack struct{
	trainedMaps train.TrainedMaps
}

func NewCrack(list int) *Crack{
	c := new(Crack)

	err := Load("maps/xaaTrainedMaps.gob", &c.tm)
	Check(err)

	return c
}

func (c Crack) guess(cmName string) chan string{
	cm := c.trainedMaps.CompositeModelsMap[cmName]

	relFreq := c.trainedMaps.RelativeFreq(cm)
	bufSize := int(MinBufferSize*relFreq)
	guessChan := make(chan string, bufSize)

	go func(){
		
	}()

	return guessChan
}

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// checks for error
func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}