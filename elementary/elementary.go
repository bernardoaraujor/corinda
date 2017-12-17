package elementary

import (
	"math"
)

// this struct represents an Elementary Model
// a map[string]ElementaryModel is later saved into a gob file
type Model struct {
	Name       string
	Entropy    float64
	TokenFreqs map[string]int
}

type TokenFreq struct {
	Token string
	Freq  int
}

func (em *Model) UpdateEntropy(){

	// sum for normalization (frequencies to probabilities)
	sum := 0
	for _, freq := range em.TokenFreqs {
		sum += freq
	}

	// entropy calculation
	entropy := float64(0)
	for _, freq := range em.TokenFreqs {
		f := freq
		p := float64(f)/float64(sum)
		e := -p*math.Log10(p)
		entropy += e
	}

	em.Entropy = entropy
}

// updates the frequency of a token in some elementary.Model
func (em *Model) UpdateTokenFreq(freq int, token string){

	//token already in TokenFreqs?
	if _, ok := em.TokenFreqs[token]; ok{
		f := em.TokenFreqs[token]
		em.TokenFreqs[token] = freq + f
	}else{
		em.TokenFreqs[token] = freq
	}
}