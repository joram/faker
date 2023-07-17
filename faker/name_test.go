package faker

import "testing"

func TestGuessSexIsMale(t *testing.T) {
	type test struct {
		name   string
		isMale bool
	}
	tests := []test{
		{"John", true},
		{"Tom", true},
		{"Robert", true},
		{"Mary", false},
		{"Carol", false},
		{"Caitlin", false},
	}
	for _, test := range tests {
		result := guessSexIsMale(test.name, 1)
		if result != test.isMale {
			t.Errorf("Expected %t but got %t\n", test.isMale, result)
		}
	}
}

func TestGetFirstName(t *testing.T) {
	type test struct {
		original string
		seed     int64
		expected string
	}
	tests := []test{
		{"John", 0, "Shawn"},
		{"Tom", 0, "William"},
		{"Robert", 0, "Jayden"},
		{"Mary", 0, "Rachel"},
		{"Carol", 0, "Ysabella"},
		{"Caitlin", 0, "Gabrielle"},

		{"John", 1, "Maxwell"},
		{"Tom", 1, "Samuel"},
		{"Robert", 1, "Marc-antoine"},
		{"Mary", 1, "Rachelle"},
		{"Carol", 1, "Deborah"},
		{"Caitlin", 1, "Kaitlyn"},
	}
	for _, test := range tests {
		result := GetFirstName(test.original, test.seed)
		if result != test.expected {
			t.Errorf("Expected %s but got %s\n", test.expected, result)
		}
	}
}

func TestGetLastName(t *testing.T) {
	type test struct {
		original string
		seed     int64
		expected string
	}
	tests := []test{
		{"John", 0, "Williams"},
		{"Tom", 0, "Desko"},
		{"Robert", 0, "Smith"},
		{"Mary", 0, "Pope"},
		{"Carol", 0, "Whaley"},
		{"Caitlin", 0, "Earnest"},

		{"John", 1, "Sheldon"},
		{"Tom", 1, "Irick"},
		{"Robert", 1, "Barnes"},
		{"Mary", 1, "Honeycutt"},
		{"Carol", 1, "Bamberger"},
		{"Caitlin", 1, "Thompson"},
	}
	for _, test := range tests {
		result := GetLastName(test.original, test.seed)
		if result != test.expected {
			t.Errorf("Expected %s but got %s\n", test.expected, result)
		}
	}
}
