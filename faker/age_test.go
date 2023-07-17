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
		expectedOutput string
	}
	tests := []test{
		{
			seed:           0,
			expectedOutput: "1989-03-12 06:59:12 -0800 PST",
		},
		{
			seed:           1,
			expectedOutput: "2009-09-16 20:32:31 -0700 PDT",
		},
	}
	for _, test := range tests {
		birthday := GetBirthday(test.seed)
		if birthday.String() != test.expectedOutput {
			t.Errorf("Expected birthday to be %s but got %s\n", test.expectedOutput, birthday)
		}
	}
}
