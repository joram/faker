package faker

import "testing"

func TestSetSeed(t *testing.T) {
	globalSeed = 0
	SetSeed(1)
	if globalSeed != 1 {
		t.Errorf("Expected globalSeed to be 1, got %d", globalSeed)
	}
}

func TestCalculateSeed(t *testing.T) {
	seed := calculateSeed("test", 1)
	expected := int64(-6441359348440544597)
	if seed != expected {
		t.Errorf("Expected seed to be %d, got %d", expected, seed)
	}
}
