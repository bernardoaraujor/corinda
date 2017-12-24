package composite

import (
	"github.com/bernardoaraujor/corinda/elementary"
)

// this struct represents a Composite Model
// a map[string]CompositeModel is later saved into a gob file
type Model struct{
	Name             string
	Freq    int
	Prob 	float64
	Entropy float64
	Models  []string
}

// gets Probability from normalized Frequency
func (cm *Model) UpdateProb(freqSum int){
	cm.Prob = float64(cm.Freq)/float64(freqSum)
}

// updates the frequency of the Composite Model
func (cm *Model) UpdateFreq(freq int){
	cm.Freq = cm.Freq + freq
}

// updates the total entropy of the Composite Model
func (cm *Model) UpdateEntropy(elementaries map[string]*elementary.Model){
	entropy := float64(0)

	for _, em := range cm.Models {
		entropy += elementaries[em].Entropy
	}

	cm.Entropy = entropy
}

// returns channel with password guesses belonging to the cartesian product between the Composite Model's Elementary Models
func (cm *Model) Guess(tokenLists [][]string) chan string{
	out := make(chan string)

	go cm.recursive(0, nil, nil, out, tokenLists)

	return out
}

// sends elements of the cartesian product of TokensNFreqs of all cm's ems to out channel
func (cm *Model) recursive(depth int, counters []int, lengths []int, out chan string, tokenLists [][]string){

	// max depth to be processed recursively
	n := len(cm.Models)

	// first depth level of recursion... init counters and lengths
	if depth == 0{

		// init counters (all 0)
		counters = make([]int, n)

		// init lengths
		lengths = make([]int, n)
		for i, _ := range cm.Models {
			//emName := cm.Models[i]
			lengths[i] = len(tokenLists[i])
		}
	}

	// last depth level of recursion
	if depth == n{
		result := ""
		for d := 0; d < n; d++{
			i := counters[d]

			result += tokenLists[d][i]
		}

		// send result to out channel
		out <- result

		// any other depth that isn't the last
	}else{
		// go through current depth
		for counters[depth] = 0; counters[depth] < lengths[depth]; counters[depth]++{
			// recursively process next depth
			cm.recursive(depth+1, counters, lengths, out, tokenLists)
		}
	}

	// time to close the channel?
	// analyzes counters of all EMs... if all are equal to the respective lengths,
	// then every element of the cartesian product have been calculated, and the channel can be closed
	closer := true
	for i := 0; i < n; i++{
		if counters[i] != lengths[i]{
			closer = false
		}
	}

	// close channel
	if closer{
		close(out)
	}

}