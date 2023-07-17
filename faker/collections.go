package faker

import (
	"github.com/joram/faker/faker/data"
	"github.com/joram/faker/faker/data_structures"
)

var (
	femaleFirstNames *data_structures.WeightedCollection
	maleFirstNames   *data_structures.WeightedCollection
	lastNames        *data_structures.WeightedCollection
	age              *data_structures.WeightedIntCollection
	maleCountBTree   *data_structures.BTree
	femaleCountBTree *data_structures.BTree
)

func getFemaleFirstNames() *data_structures.WeightedCollection {
	if femaleFirstNames == nil {
		femaleFirstNames = data_structures.NewWeightedCollection(&data.FemaleNames)
	}
	return femaleFirstNames
}

func getMaleFirstNames() *data_structures.WeightedCollection {
	if maleFirstNames == nil {
		maleFirstNames = data_structures.NewWeightedCollection(&data.MaleNames)
	}
	return maleFirstNames
}

func getLastNames() *data_structures.WeightedCollection {
	if lastNames == nil {
		lastNames = data_structures.NewWeightedCollection(&data.LastNames)
	}
	return lastNames
}

func getAge() *data_structures.WeightedIntCollection {
	if age == nil {
		age = data_structures.NewWeightedIntCollection(&data.Ages)
	}
	return age
}

func getMaleCountBTree() *data_structures.BTree {
	if maleCountBTree == nil {
		maleCountBTree = data_structures.NewBTree(data.MaleNames)
	}
	return maleCountBTree
}

func getFemaleCountBTree() *data_structures.BTree {
	if femaleCountBTree == nil {
		femaleCountBTree = data_structures.NewBTree(data.FemaleNames)
	}
	return femaleCountBTree
}
