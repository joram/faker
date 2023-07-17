package faker

func GetMaleFirstName(originalName string, seed int64) string {
	return getMaleFirstNames().GetRandomName(
		calculateSeed(originalName, seed),
	)
}

func GetFemaleFirstName(originalName string, seed int64) string {
	return getFemaleFirstNames().GetRandomName(
		calculateSeed(originalName, seed),
	)
}

func GetFirstName(originalName string, seed int64) string {
	seed = calculateSeed(originalName, seed)
	if guessSexIsMale(originalName, seed) {
		return GetMaleFirstName(originalName, seed)
	}
	return GetFemaleFirstName(originalName, seed)
}

func GetLastName(originalName string, seed int64) string {
	return getLastNames().GetRandomName(
		calculateSeed(originalName, seed),
	)
}
