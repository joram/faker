package faker

import (
	"math/rand"
)

func guessSexIsMale(originalName string, seed int64) bool {
	numMales, _ := getMaleCountBTree().Get(originalName)
	numFemales, _ := getFemaleCountBTree().Get(originalName)
	total := numMales + numFemales
	if total == 0 {
		return rand.Intn(2) == 0
	}
	randNum := rand.Intn(total)
	chooseMale := randNum < numMales
	return chooseMale
}
