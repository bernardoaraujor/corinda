package elementary

import (
	"sort"
	"github.com/sajari/regression"
	"math"
)

// this struct represents an Elementary Model
// a map[string]ElementaryModel is later saved into a gob file
type Model struct {
	Name         string
	Complexity   int
	Entropy      float64
	S            float64
	ZetaInv      float64
	TokensNfreqs []TokenNfreq
}

type TokenNfreq struct {
	Token string
	Freq  int
}

func (em *Model) UpdateEntropy(){
	// make sure frequencies are sorted
	em.Sort()

	// linear regression in log-log (Freqs follow a zeta distribution)
	r := new(regression.Regression)
	for k, tf := range em.TokensNfreqs{
		logK := math.Log10(float64(k+1))
		logF := math.Log10(float64(tf.Freq))

		dp := regression.DataPoint(logF, []float64{logK})
		r.Train(dp)

	}
	r.Run()

	// these are the coefficients of the linear regression from log-log
	b := r.Coeff(0)
	a := r.Coeff(1)

	// zeta parames
	em.S = -a
	em.ZetaInv = math.Pow(10, b)

	// entropy = sum of -P*log10P
	entropy := float64(0)
	for k := 1; k <= em.Complexity; k++{
		p := math.Pow(float64(k), -em.S)
		entropy -= p*math.Log10(p)
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
		em.TokensNfreqs = append(em.TokensNfreqs, TokenNfreq{token, freq})
	}

	em.Sort()
}

// sorts TokensNfreqs in descendent order
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
