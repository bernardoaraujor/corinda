package composite

import "github.com/bernardoaraujor/corinda/elementary"

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