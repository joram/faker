package faker

import "time"

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
	Birthday   time.Time
	Email      string
	Phone      string
	Address    Address
	CreditCard CreditCard
}

func GetPerson(seed int64) Person {
	firstName := GetFirstName("", seed)
	middleName := GetFirstName(firstName, seed)

	return Person{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   GetLastName("", seed),
		Birthday:   GetBirthday(seed),
		Email:      GetEmail("", seed),
		Phone:      GetPhone("", seed),
		Address:    GetAddress(seed),
		CreditCard: GetCreditCard(seed),
	}
}
