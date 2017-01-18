package mapkha

import (
	"reflect"
	"testing"
)

func TestOneCharPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"A", 10}}
	prefixTree := MakePrefixTree(words)
	expect := &PrefixTreePointer{0, true, 10}
	testLookup(t, expect, "Expect to find 0, 0, A")(prefixTree.Lookup(0, 0, 'A'))
}

func TestOneWordPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"AB", 20}}
	prefixTree := MakePrefixTree(words)

	var expect *PrefixTreePointer

	expect = &PrefixTreePointer{0, false, nil}
	testLookup(t, expect, "Expect to find 0, 0, A")(prefixTree.Lookup(0, 0, 'A'))

	expect = &PrefixTreePointer{0, true, 20}
	testLookup(t, expect, "Expect to find 0, 1, B")(prefixTree.Lookup(0, 1, 'B'))
}

func TestTwoWordsPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"AB", 20}, {"AC", 30}, {"D", 40}}
	prefixTree := MakePrefixTree(words)

	var expect *PrefixTreePointer

	expect = &PrefixTreePointer{0, false, nil}
	testLookup(t, expect, "Expect to find 0, 0, A")(prefixTree.Lookup(0, 0, 'A'))

	expect = &PrefixTreePointer{0, true, 20}
	testLookup(t, expect, "Expect to find 0, 1, B")(prefixTree.Lookup(0, 1, 'B'))

	expect = &PrefixTreePointer{1, true, 30}
	testLookup(t, expect, "Expect to find 0, 1, C")(prefixTree.Lookup(0, 1, 'C'))

	expect = &PrefixTreePointer{2, true, 40}
	testLookup(t, expect, "Expect to find 0, 0, D")(prefixTree.Lookup(0, 0, 'D'))
}

func testLookup(t *testing.T, expect *PrefixTreePointer, msg string) func(*PrefixTreePointer, bool) {
	return func(child *PrefixTreePointer, found bool) {
		if !found {
			t.Errorf(msg)
		}

		if !reflect.DeepEqual(expect, child) {
			t.Errorf("Expect %q got %q", expect, child)
		}
	}
}
