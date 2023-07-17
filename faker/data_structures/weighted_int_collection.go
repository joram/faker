package data_structures

import (
	"math/rand"
	"sort"
)

type WeightedIntCollection struct {
	values  []int
	indices []int
}

func (wg *WeightedIntCollection) GetRandomValue(seed int64) int {
	rand.Seed(seed)
	i := rand.Intn(len(wg.indices))
	index := wg.indices[i]
	return wg.values[index]
}

func NewWeightedIntCollection(data *map[int]int) *WeightedIntCollection {
	wc := &WeightedIntCollection{
		values:  []int{},
		indices: []int{},
	}

	type nameChance struct {
		name       int
		chances    int
		percentage float64
	}

	names := []int{}
	for name, _ := range *data {
		names = append(names, name)
	}
	sort.Ints(names)

	nameChances := []nameChance{}
	setMinChance := false
	minChance := 0
	for _, name := range names {
		chance := (*data)[name]
		nameChances = append(nameChances, nameChance{
			name:    name,
			chances: chance,
		})

		if !setMinChance || chance < minChance {
			minChance = chance
			setMinChance = true
		}
	}

	for _, nameChance := range nameChances {
		wc.values = append(wc.values, nameChance.name)
		nameIndex := len(wc.values) - 1
		for i := 0; i < nameChance.chances; i++ {
			wc.indices = append(wc.indices, nameIndex)
		}
	}
	return wc
}
