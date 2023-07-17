package faker

import (
	"github.com/joram/faker/faker/data"
	"math/rand"
)

type Address struct {
	Country      string
	State        string
	City         string
	StreetNumber string
	Street       string
	ZipCode      string
}

func GetAddress(seed int64) Address {
	rand.Seed(seed)
	i := rand.Intn(len(data.Locations) - 1)
	location := data.Locations[i]
	if location.Province == "Newfoundland and Labrador" {
		location.Province = []string{"Newfoundland", "Labrador"}[rand.Intn(2)]
	}

	i = rand.Intn(len(location.PostalCodes) - 1)
	postalCode := location.PostalCodes[i]

	return Address{
		Country:      "CA",
		State:        location.Province,
		City:         location.City,
		StreetNumber: getRandomNumber("", 4),
		Street:       "fake street",
		ZipCode:      postalCode,
	}
}
