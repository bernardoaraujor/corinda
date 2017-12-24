package elementary

import (
	"math"
	"sort"
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

// temporary structure used to sort tokens
type tokenSort struct{
	tokenFreqs map[string]int
	tokenSlice []string
}

//returns a sorted slice with tokens in descending frequency order
func (em *Model) SortedTokens() []string{
	tokenSlice := make([]string, 0)

	for token, _ := range em.TokenFreqs{
		tokenSlice = append(tokenSlice, token)
	}

	tokenSort := tokenSort{em.TokenFreqs, tokenSlice}
	tokenSort.Sort()

	return tokenSort.tokenSlice
}

// sorts TokenFreqs in descendent order
func (ts *tokenSort) Sort(){
	sort.Sort(ts)
}

// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic.
func (ts *tokenSort) Len() int {
	return len(ts.tokenFreqs)
}

func (ts *tokenSort) Swap(i, j int) {
	ts.tokenSlice[i], ts.tokenSlice[j] = ts.tokenSlice[j], ts.tokenSlice[i]
}
func (ts *tokenSort) Less(i, j int) bool {
	tokenI := ts.tokenSlice[i]
	tokenJ := ts.tokenSlice[j]

	return ts.tokenFreqs[tokenI] > ts.tokenFreqs[tokenJ]
}