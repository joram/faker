package faker

import (
	"fmt"
	"math/rand"
)

func GetPhone(originalPhone string, seed int64) string {
	seed = calculateSeed(originalPhone, seed)
	rand.Seed(seed)

	countryCode := 1
	areaCode := getRandomNumber("", 3)
	phone1 := getRandomNumber("", 3)
	phone2 := getRandomNumber("", 4)
	return fmt.Sprintf("%d-%s-%s-%s", countryCode, areaCode, phone1, phone2)
}
