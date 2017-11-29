package composite

import (
	"github.com/bernardoaraujor/corinda/elementary"
	"math/rand"
	"time"
)

// this struct represents a Composite Model
// a map[string]CompositeModel is later saved into a gob file
type Model struct{
	Name             string
	Complexity       int
	Freq             int
	Entropy float64
	ElementaryModels []*elementary.Model
}

// updates the frequency of the Composite Model
func (cm *Model) UpdateFreq(freq int){
	cm.Freq = cm.Freq + freq
}

// updates the total entropy of the Composite Model
func (cm *Model) UpdateEntropy(){
	entropy := float64(0)

	for _, em := range cm.ElementaryModels{
		entropy += em.Entropy
	}

	cm.Entropy = entropy
}

// returns channel with password guesses belonging to the cartesian product between the Composite Model's Elementary Models
func (cm *Model) Guess() chan string{
	out := make(chan string)

	go cm.recursive(0, nil, nil, out)

	return out
}

// sends elements of the cartesian product of TokensNFreqs of all cm's ems to out channel
func (cm *Model) recursive(depth int, counters []int, lengths []int, out chan string){
	// max depth to be processed recursively
	n := len(cm.ElementaryModels)

	// first depth level of recursion... init counters and lengths
	if depth == 0{

		// init counters (all 0)
		counters = make([]int, n)

		// init lengths
		lengths = make([]int, n)
		for i, _ := range cm.ElementaryModels{
			lengths[i] = len(cm.ElementaryModels[i].TokensNfreqs)
		}
	}

	// last depth level of recursion
	if depth == n{
		resultado := ""
		for d := 0; d < n; d++{
			i := counters[d]
			resultado += cm.ElementaryModels[d].TokensNfreqs[i].Token
		}

		// send result to out channel
		out <- resultado

		// any other depth that isn't the last
	}else{
		// sweep current depth
		for counters[depth] = 0; counters[depth] < lengths[depth]; counters[depth]++{
			// recursively process next depth
			cm.recursive(depth+1, counters, lengths, out)
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

// randomly sends combinations of tokens to out channel
func (cm *Model) randomGuess() chan string{
	out := make(chan string)

	go func(){
		s1 := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s1)
		for {
			guess := ""
			for _, em := range cm.ElementaryModels{
				i := r.Intn(len(em.TokensNfreqs))
				guess += em.TokensNfreqs[i].Token
			}

			out <- guess
		}
	}()

	return out
}