package faker

import (
	"testing"
)

func TestMarkovChainGetRandomWord(t *testing.T) {
	type test struct {
		seed     int64
		expected string
	}

	tests := []test{
		{0, "intelligenced"},
		{1, "eradiates"},
	}

	for _, test := range tests {
		word := getEnglishWordsMarkovChain().GetRandomWord(test.seed)
		if word != test.expected {
			t.Errorf("expected %s, but got %s\n", test.expected, word)
		}
	}
}
