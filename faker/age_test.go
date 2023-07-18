package faker

import "testing"

func TestGetAge(t *testing.T) {
	type test struct {
		seed        int64
		expectedAge int
	}
	tests := []test{
		{
			seed:        0,
			expectedAge: 34,
		},
		{
			seed:        1,
			expectedAge: 14,
		},
	}
	for _, test := range tests {
		age := GetAge(test.seed)
		if age != test.expectedAge {
			t.Errorf("Expected age to be %d but got %d\n", test.expectedAge, age)
		}
	}
}

func TestGetBirthday(t *testing.T) {
	type test struct {
		seed           int64
		expectedOutput int64
	}
	tests := []test{
		{
			seed:           0,
			expectedOutput: 605717952,
		},
		{
			seed:           1,
			expectedOutput: 1253158351,
		},
	}
	for _, test := range tests {
		birthday := GetBirthday(test.seed)
		if birthday.Unix() != test.expectedOutput {
			t.Errorf("Expected birthday to be %d but got %.d\n", test.expectedOutput, birthday.Unix())
		}
	}
}
