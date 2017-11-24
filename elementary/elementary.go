package elementary

import (
	"sort"
	"github.com/sajari/regression"
	"math"
	//"fmt"
	//"strconv"
)

// this struct represents an Elementary Model
// a map[string]ElementaryModel is later saved into a gob file
type Model struct {
	Name         string
	Complexity   int
	Entropy      float64
	TokensNfreqs []TokenNfreq
}

type TokenNfreq struct {
	Token string
	Freq  int
}

func (em *Model) UpdateEntropy(){
	// make sure frequencies are sorted
	em.Sort()

	// linear regression in log-log (long tail is a power law)
	// using graphical method... TODO: maximum likelihood would be better?
	r := new(regression.Regression)
	kmin := int(float64(len(em.TokensNfreqs))*0.8)		// long tail is the last 20%
	kmax := len(em.TokensNfreqs)
	for k := kmin; k < kmax; k++{
		logK := math.Log10(float64(k+1))
		logF := math.Log10(float64(em.TokensNfreqs[k].Freq))

		dp := regression.DataPoint(logF, []float64{logK})
		r.Train(dp)
	}
	r.Run()

	// these are the coefficients of the linear regression from log-log
	a := r.Coeff(1)
	b := r.Coeff(0)

	alpha := -a
	c := math.Pow(10, b)

	// generate virtual freqs
	maxSize := 10000000
	size := em.Complexity
	if size > maxSize{		// we don't want to run out of memory! besides, after this limit only brute force models are affected
		size = maxSize
	}

	vFreqs := make([]float64, size)
	sum := float64(0)
	for k, _ := range vFreqs{
		if k < kmax{		// use actual freqs
			f := float64(em.TokensNfreqs[k].Freq)
			vFreqs[k] = f
			sum += f
		}else{
			f := c*math.Pow(float64(k), -alpha)
			vFreqs[k] = f
			sum += f
		}
	}

	// calc entropy
	entropy := float64(0)
	for _, vf := range vFreqs{
		p := vf/sum
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
