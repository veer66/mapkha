package mapkha

import (
	"reflect"
	"testing"
)

func TestOneCharPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"A", 10}}
	prefixTree := MakePrefixTree(words)
	expect := &PrefixTreePointer{0, true, 10}
	child, found := prefixTree.Lookup(0, 0, 'A')

	if !found {
		t.Errorf("Expect to find 0, 0, A")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}
}

func TestOneWordPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"AB", 20}}
	prefixTree := MakePrefixTree(words)

	var expect *PrefixTreePointer
	var child *PrefixTreePointer
	var found bool

	expect = &PrefixTreePointer{0, false, nil}
	child, found = prefixTree.Lookup(0, 0, 'A')

	if !found {
		t.Errorf("Expect to find 0, 0, A")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}

	expect = &PrefixTreePointer{0, true, 20}
	child, found = prefixTree.Lookup(0, 1, 'B')

	if !found {
		t.Errorf("Expect to find 0, 1, B")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}
}

func TestTwoWordsPrefixTree(t *testing.T) {
	words := []WordWithPayload{{"AB", 20}, {"AC", 30}, {"D", 40}}
	prefixTree := MakePrefixTree(words)

	var expect *PrefixTreePointer
	var child *PrefixTreePointer
	var found bool

	expect = &PrefixTreePointer{0, false, nil}
	child, found = prefixTree.Lookup(0, 0, 'A')

	if !found {
		t.Errorf("Expect to find 0, 0, A")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}

	expect = &PrefixTreePointer{0, true, 20}
	child, found = prefixTree.Lookup(0, 1, 'B')

	if !found {
		t.Errorf("Expect to find 0, 1, B")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}

	expect = &PrefixTreePointer{1, true, 30}
	child, found = prefixTree.Lookup(0, 1, 'C')

	if !found {
		t.Errorf("Expect to find 0, 1, C")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}

	expect = &PrefixTreePointer{2, true, 40}
	child, found = prefixTree.Lookup(0, 0, 'D')

	if !found {
		t.Errorf("Expect to find 0, 0, D")
	}

	if !reflect.DeepEqual(expect, child) {
		t.Errorf("Expect %q got %q", expect, child)
	}

}
