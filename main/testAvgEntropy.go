// MEDIA PONDERADA DA ENTROPIA!!

package main

import (
	"os"
	"compress/gzip"
	"io/ioutil"
	"github.com/bernardoaraujor/corinda/composite"
	"encoding/json"
	"runtime"
	"fmt"
	"strconv"
)

func main() {

	f, err := os.Open("maps/rockyouComposites.json.gz")
	check(err)
	defer f.Close()
	gr, err := gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cm, err := ioutil.ReadAll(gr)
	check(err)
	var composites1 []*composite.Model
	err = json.Unmarshal(cm, &composites1)
	check(err)

	prob, err := os.Create("rockyou_prob.txt")
	entr, err := os.Create("rockyou_entr.txt")

	sumEntropy := 0.0
	sumProb := 0.0
	for _, composite := range composites1{
		sumEntropy += composite.Prob*composite.Entropy
		sumProb += composite.Prob

		fmt.Fprintln(prob, strconv.FormatFloat(composite.Prob, 'f', -1, 64))
		fmt.Fprintln(entr, composite.Entropy)
	}

	avgEntropy := sumEntropy
	avgProb := sumProb/float64(len(composites1))
	fmt.Println("rockyou avg entropy: " + strconv.FormatFloat(avgEntropy, 'f', 10, 64) + ", avg prob: " + strconv.FormatFloat(avgProb, 'f', 10, 64))

	f, err = os.Open("maps/linkedin_1MComposites.json.gz")
	check(err)
	defer f.Close()
	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cm, err = ioutil.ReadAll(gr)
	check(err)
	var composites2 []*composite.Model
	err = json.Unmarshal(cm, &composites2)
	check(err)

	prob, err = os.Create("linkedin_1M_prob.txt")
	entr, err = os.Create("linkedin_1M_entr.txt")

	sumEntropy = 0.0
	sumProb = 0.0
	for _, composite := range composites2{
		sumEntropy += composite.Entropy
		sumProb += composite.Prob

		fmt.Fprintln(prob, strconv.FormatFloat(composite.Prob, 'f', -1, 64))
		fmt.Fprintln(entr, composite.Entropy)
	}

	avgEntropy = sumEntropy/float64(len(composites2))
	avgProb = sumProb/float64(len(composites2))
	fmt.Println("linkedin avg entropy: " + strconv.FormatFloat(avgEntropy, 'f', 10, 64) + ", avg prob: " + strconv.FormatFloat(avgProb, 'f', 10, 64))

	f, err = os.Open("maps/antipublic_1MComposites.json.gz")
	check(err)
	defer f.Close()
	gr, err = gzip.NewReader(f)
	check(err)
	defer gr.Close()

	cm, err = ioutil.ReadAll(gr)
	check(err)
	var composites3 []*composite.Model
	err = json.Unmarshal(cm, &composites3)
	check(err)

	prob, err = os.Create("antipublic_1M_prob.txt")
	entr, err = os.Create("antipublic_1M_entr.txt")

	sumEntropy = 0.0
	sumProb = 0.0
	for _, composite := range composites3{
		sumEntropy += composite.Entropy
		sumProb += composite.Prob

		fmt.Fprintln(prob, strconv.FormatFloat(composite.Prob, 'f', -1, 64))
		fmt.Fprintln(entr, composite.Entropy)
	}

	avgEntropy = sumEntropy/float64(len(composites2))
	avgProb = sumProb/float64(len(composites2))
	fmt.Println("antipublic avg entropy: " + strconv.FormatFloat(avgEntropy, 'f', 10, 64) + ", avg prob: " + strconv.FormatFloat(avgProb, 'f', 10, 64))

}

// checks for error
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
