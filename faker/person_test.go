package faker

import (
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

func TestGetPerson(t *testing.T) {
	person := GetPerson(0)
	spew.Dump(person)

	expected := Person{
		FirstName:  "Radin",
		MiddleName: "Laurier",
		LastName:   "Timko",
		Birthday:   GetBirthday(0),
		Email:      "grouses@versicles.ca",
		Phone:      "1-319-772-5167",
		Address: Address{
			Country:      "CA",
			State:        "Newfoundland",
			City:         "Happy Valley-goose Bay",
			StreetNumber: "6567",
			Street:       "fake street",
			ZipCode:      "A0P1S0",
		},
		CreditCard: CreditCard{
			CardType:    0,
			Number:      4716888798261,
			ExpiryMonth: 12,
			ExpiryYear:  2029,
			CVV:         443,
		},
	}

	if !reflect.DeepEqual(expected, person) {

		t.Errorf("Expected %#v,\ngot %#v", expected, person)
	}
}
