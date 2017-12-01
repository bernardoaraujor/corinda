package main

import (
	"os"
	"runtime"
	"fmt"
	"encoding/csv"
	"crypto/sha1"
	"encoding/hex"
)

func main() {
	list := "rockyou"
	f, err := os.Open("csv/" + list + ".csv")
	Check(err)
	defer f.Close()

	cr := csv.NewReader(f)

	result, err := os.Create("targets/sha1/" + list + ".csv")
	Check(err)
	defer result.Close()

	for record, err := cr.Read(); record != nil; record, err = cr.Read(){
		Check(err)

		pass := record[1]

		hasher := sha1.New()
		hasher.Write([]byte(pass))
		digest := hasher.Sum(nil)
		hash := hex.EncodeToString(digest)

		fmt.Fprintln(result, hash)
	}
}

// checks for error
func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}