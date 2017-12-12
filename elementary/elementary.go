package elementary

import (
	"sort"
	"math"
)

// this struct represents an Elementary Model
// a map[string]ElementaryModel is later saved into a gob file
type Model struct {
	Name       string
	Entropy    float64
	TokenFreqs []TokenFreq
}

type TokenFreq struct {
	Token string
	Freq  int
}

func (em *Model) UpdateEntropy(){

	// sum for normalization (frequencies to probabilities)
	sum := 0
	for _, tf := range em.TokenFreqs {
		sum += tf.Freq
	}

	// entropy calculation
	entropy := float64(0)
	for _, tf := range em.TokenFreqs {
		f := tf.Freq
		p := float64(f)/float64(sum)
		e := -p*math.Log10(p)
		entropy += e
	}

	em.Entropy = entropy
}

// updates the frequency of a token in some ElementaryModel
func (em *Model) UpdateTokenFreq(freq int, token string){
	// is token already in em?
	index := 0
	b := false
	for i, tf := range em.TokenFreqs {
		if tf.Token == token{
			index = i
			b = true
			break
		}
	}

	if b{		// yes, token is in em
		em.TokenFreqs[index].Freq += freq
	}else{		//no, token is not in em
		em.TokenFreqs = append(em.TokenFreqs, TokenFreq{token, freq})
	}

	em.Sort()
}

// sorts TokenFreqs in descendent order
func (em *Model) Sort(){
	sort.Sort(em)
}

// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic.
func (em *Model) Len() int {
	return len(em.TokenFreqs)
}
func (em *Model) Swap(i, j int) {
	em.TokenFreqs[i], em.TokenFreqs[j] = em.TokenFreqs[j], em.TokenFreqs[i]
}
func (em *Model) Less(i, j int) bool {
	return em.TokenFreqs[i].Freq > em.TokenFreqs[j].Freq
}
