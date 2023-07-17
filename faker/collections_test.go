package faker

import "testing"

func TestGetFemaleFirstName(t *testing.T) {
	femaleFirstNames = nil
	femaleFirstNames = getFemaleFirstNames()
	if femaleFirstNames == nil {
		t.Error("Expected femaleFirstNames to be initialized")
	}
}

func TestGetMaleFirstName(t *testing.T) {
	maleFirstNames = nil
	maleFirstNames = getMaleFirstNames()
	if maleFirstNames == nil {
		t.Error("Expected maleFirstNames to be initialized")
	}
}

func Test_getLastName(t *testing.T) {
	lastNames = nil
	lastNames = getLastNames()
	if lastNames == nil {
		t.Error("Expected lastNames to be initialized")
	}
}

func TestGetMaleCountBTree(t *testing.T) {
	maleCountBTree = nil
	maleCountBTree = getMaleCountBTree()
	if maleCountBTree == nil {
		t.Error("Expected maleCountBTree to be initialized")
	}
}

func TestGetFemaleCountBTree(t *testing.T) {
	femaleCountBTree = nil
	femaleCountBTree = getFemaleCountBTree()
	if femaleCountBTree == nil {
		t.Error("Expected femaleCountBTree to be initialized")
	}
}
