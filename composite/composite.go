package composite

import "github.com/bernardoaraujor/corinda/elementary"

// this struct represents a Composite Model
// a map[string]CompositeModel is later saved into a gob file
type Model struct{
	Name             string
	Complexity       int
	Freq             int
	ElementaryModels []*elementary.Model
}

// updates the frequency of some CompositeModel
func (cm *Model) UpdateFreq(freq int){
	cm.Freq = cm.Freq + freq
}
