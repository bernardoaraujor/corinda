package crack

import (
	"runtime"
	"fmt"
	"os"
	"github.com/bernardoaraujor/corinda/train"
	"encoding/gob"
	"crypto/sha1"
	"hash"
	"crypto/sha256"
)

const rockyou = "rockyou"
const linkedin = "linkedin"
const antipublic = "antipublic"

const SHA1 = "SHA1"
const SHA256 = "SHA256"

type Crack struct {
	trainedMaps train.TrainedMaps
	alg         string
}

// Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

// returns channel with hashes in string format, and iterates n times retrieving password guesses over in
// digest implements the fan in patterns
func (crack Crack) Digest(in chan string, n int) chan []byte {
	out := make(chan []byte)

	go func(n int, out chan []byte) {
		defer close(out)
		for i := 0; i < n; i++ {
			// reads in channel
			s := <-in

			// digest
			var hasher hash.Hash
			switch crack.alg{
			case SHA1:
				hasher = sha1.New()
			case SHA256:
				hasher = sha256.New()
			}
			hasher.Write([]byte(s))
			digest := hasher.Sum(nil)

			// spits out digest
			out <- digest
		}
	}(n, out)

	return out
}
