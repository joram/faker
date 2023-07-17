package faker

import (
	"fmt"
	"github.com/tjarratt/babble"
	"math/rand"
	"strings"
)

var (
	TLDs = []string{"com", "net", "org", "ca"}
)

func GetEmail(originalEmail string, seed int64) string {
	seed = calculateSeed(originalEmail, seed)
	rand.Seed(seed)

	babbler := babble.NewBabbler()
	babbler.Separator = "_"
	babbler.Count = 1

	username := babbler.Babble()
	babbler.Count = 2
	domain := babbler.Babble()
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
