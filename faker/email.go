package faker

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	TLDs = []string{"com", "net", "org", "ca"}
)

func GetEmail(originalEmail string, seed int64) string {
	seed = calculateSeed(originalEmail, seed)
	rand.Seed(seed)

	username := getEnglishWordsMarkovChain().GetRandomWord(seed)
	domain := getEnglishWordsMarkovChain().GetRandomWord(seed + 1)

	tld := TLDs[rand.Intn(len(TLDs))]
	email := fmt.Sprintf("%s@%s.%s", username, domain, tld)
	email = sanitizeEmail(email)

	return email
}

func sanitizeEmail(email string) string {
	for _, s := range []string{" ", "'", "\""} {
		email = strings.Replace(email, s, "", -1)
	}
	email = strings.ToLower(email)

	return email
}
