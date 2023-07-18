package data_structures

import (
	"math/rand"
)

type MarkovChain struct {
	subChains map[string]*MarkovChain
	weights   *WeightedCollection
}

func (mc *MarkovChain) addWord(word string) {
	if len(word) == 0 {
		return
	}

	c := string(word[0])
	if mc.subChains == nil {
		mc.subChains = map[string]*MarkovChain{}
	}
	if mc.weights == nil {
		mc.weights = &WeightedCollection{}
	}
	mc.weights.Add(c)
	if mc.subChains[c] == nil {
		mc.subChains[c] = &MarkovChain{}
	}
	mc.subChains[c].addWord(word[1:])
}

func (mc *MarkovChain) GetRandomWord(seed int64) string {
	rand.Seed(seed)
	if mc.weights == nil {
		return ""
	}
	c := mc.weights.GetRandomName(seed)
	if mc.subChains[c] == nil {
		panic("subchain is nil")
	}
	return c + mc.subChains[c].GetRandomWord(seed)
}

func NewMarkovChain(words []string) *MarkovChain {
	mc := &MarkovChain{
		subChains: map[string]*MarkovChain{},
		weights:   &WeightedCollection{},
	}

	for _, word := range words {
		mc.addWord(word)
	}

	return mc
}
