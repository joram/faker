package data_structures

import (
	"math/rand"
	"sort"
)

type WeightedCollection struct {
	names   []string
	indices []int
}

func (wg *WeightedCollection) GetRandomName(seed int64) string {
	rand.Seed(seed)
	i := rand.Intn(len(wg.indices))
	index := wg.indices[i]
	return wg.names[index]
}

func NewWeightedCollection(data *map[string]int) *WeightedCollection {
	wc := &WeightedCollection{
		names:   []string{},
		indices: []int{},
	}

	type nameChance struct {
		name       string
		chances    int
		percentage float64
	}

	names := []string{}
	for name, _ := range *data {
		names = append(names, name)
	}
	sort.Strings(names)

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
		wc.names = append(wc.names, nameChance.name)
		nameIndex := len(wc.names) - 1
		for i := 0; i < nameChance.chances; i++ {
			wc.indices = append(wc.indices, nameIndex)
		}
	}
	return wc
}
