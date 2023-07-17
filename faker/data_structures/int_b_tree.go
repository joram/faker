package data_structures

import "errors"

var ImproperlyInitializedBooleanBTree = errors.New("improperly initialized BTree")
var StringNotFoundInBooleanBTree = errors.New("string not found in BTree")

type BTree struct {
	isLeaf bool
	value  int
	data   map[rune]*BTree
}

func (b *BTree) add(s string, i int) {
	if len(s) == 0 {
		b.isLeaf = true
		b.value = i
		return
	}

	r := rune(s[0])
	if b.data == nil {
		b.data = map[rune]*BTree{}
	}
	if b.data[r] == nil {
		b.data[r] = &BTree{}
	}
	b.data[r].add(s[1:], i)
}

func (b *BTree) Get(s string) (int, error) {
	if len(s) == 0 {
		return b.value, nil
	}

	r := rune(s[0])
	if b.data == nil {
		return 0, ImproperlyInitializedBooleanBTree
	}
	if b.data[r] == nil {
		return 0, StringNotFoundInBooleanBTree
	}

	return b.data[r].Get(s[1:])
}

func NewBTree(data map[string]int) *BTree {
	b := &BTree{}
	for name, value := range data {
		b.add(name, value)
	}
	return b
}
