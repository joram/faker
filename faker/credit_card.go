package faker

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//The first digit is different for each card network:
//
//Visa cards begin with a 4 and have 13 or 16 digits
//Mastercard cards begin with a 5 and has 16 digits
const (
	VISA = iota
	MASTERCARD
)

var (
	visaPrefixList       = []string{"4539", "4556", "4916", "4532", "4929", "40240071", "4485", "4716", "4"}
	mastercardPrefixList = []string{"51", "52", "53", "54", "55"}
	visaLengths          = []int{13, 16}
	mastercardLengths    = []int{16}
	cardTypes            = []int{VISA, MASTERCARD}
)

type CreditCard struct {
	CardType    int
	Number      int
	ExpiryMonth int
	ExpiryYear  int
	CVV         int
}

func getRandomNumber(prefix string, length int) string {
	n := prefix
	for len(n) < length {
		n += fmt.Sprint(rand.Intn(10))
	}
	return n
}

func getCreditCardNumber(prefixList []string, lengths []int) string {
	prefix := prefixList[rand.Intn(len(prefixList)-1)]
	length := lengths[rand.Intn(len(lengths)-1)]
	return getRandomNumber(prefix, length)
}

func GetCreditCard(seed int64) CreditCard {
	rand.Seed(seed)
	cvv, _ := strconv.Atoi(getRandomNumber("", 3))

	card := CreditCard{
		CardType:    cardTypes[rand.Intn(len(cardTypes)-1)],
		CVV:         cvv,
		ExpiryMonth: rand.Intn(12) + 1,
		ExpiryYear:  time.Now().Year() + rand.Intn(10),
	}

	number := ""
	if card.CardType == VISA {
		number = getCreditCardNumber(visaPrefixList, visaLengths)
	} else {
		number = getCreditCardNumber(mastercardPrefixList, mastercardLengths)
	}
	card.Number, _ = strconv.Atoi(number)

	return card
}
