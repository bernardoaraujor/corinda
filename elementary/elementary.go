package elementary

import (
	"sort"
	"math"
)

// this struct represents an Elementary Model
// a map[string]ElementaryModel is later saved into a gob file
type Model struct {
	Name         string
	Entropy      float64
	TokensNfreqs []TokenFreq
}

type TokenFreq struct {
	Token string
	Freq  int
}

func (em *Model) UpdateEntropy(){

	// sum for normalization (frequencies to probabilities)
	sum := 0
	for _, tf := range em.TokensNfreqs{
		sum += tf.Freq
	}

	// entropy calculation
	entropy := float64(0)
	for _, tf := range em.TokensNfreqs{
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
	for i, tf := range em.TokensNfreqs{
		if tf.Token == token{
			index = i
			b = true
			break
		}
	}

	if b{		// yes, token is in em
		em.TokensNfreqs[index].Freq += freq
	}else{		//no, token is not in em
		em.TokensNfreqs = append(em.TokensNfreqs, TokenFreq{token, freq})
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
	return len(em.TokensNfreqs)
}
func (em *Model) Swap(i, j int) {
	em.TokensNfreqs[i], em.TokensNfreqs[j] = em.TokensNfreqs[j], em.TokensNfreqs[i]
}
func (em *Model) Less(i, j int) bool {
	return em.TokensNfreqs[i].Freq > em.TokensNfreqs[j].Freq
}
