package faker

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

var ErrInvalidIPAddress = errors.New("improperly formatted IP address")

func GetIPAddress(original string, seed int64) (string, error) {
	n := strings.Count(original, ".")
	if n != 3 {
		return "", ErrInvalidIPAddress
	}

	parts := strings.Split(original, ".")
	if len(parts) != 4 {
		return "", ErrInvalidIPAddress
	}

	var octets []int
	for _, octet := range parts {
		rand.Seed(calculateSeed(octet, seed))
		newOctet := rand.Intn(255)
		octets = append(octets, newOctet)
	}

	newIP := fmt.Sprintf("%d.%d.%d.%d", octets[0], octets[1], octets[2], octets[3])
	return newIP, nil
}
