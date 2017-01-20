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

func TestKaPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"กา", true}}
	prefixTree := MakePrefixTree(words)
	var expect *PrefixTreePointer

	expect = &PrefixTreePointer{0, false, nil}
	testLookup(t, expect, "Expect to find 0, 0, ก")(prefixTree.Lookup(0, 0, 'ก'))

	expect = &PrefixTreePointer{0, true, true}
	testLookup(t, expect, "Expect to find 0, 1, า")(prefixTree.Lookup(0, 1, 'า'))
}

func TestViaDict(t *testing.T) {
	dict, _ := LoadDefaultDict()
	var child *PrefixTreePointer
	var found bool

	child, found = dict.Lookup(0, 0, 'ม')
	if !found {
		t.Errorf("Expect to find ม")
	}

	child, found = dict.Lookup(child.ChildID, 1, 'า')
	if !found {
		t.Errorf("Expect to find า")
	}

	child, found = dict.Lookup(child.ChildID, 2, 'ต')
	if !found {
		t.Errorf("Expect to find ต")
	}

	child, found = dict.Lookup(child.ChildID, 3, 'ร')
	if !found {
		t.Errorf("Expect to find ร")
	}

	child, found = dict.Lookup(child.ChildID, 4, 'า')
	if !found {
		t.Errorf("Expect to find last า")
	}

	if !child.IsFinal {
		t.Errorf("Expect last า to be final")
	}

}

func TestViaDictNotFinal(t *testing.T) {
	dict, _ := LoadDefaultDict()
	var child *PrefixTreePointer

	child, _ = dict.Lookup(0, 0, 'ต')

	child, _ = dict.Lookup(child.ChildID, 1, 'ร')

	if child.IsFinal {
		t.Errorf("Expect last ร not to be final")
	}

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
