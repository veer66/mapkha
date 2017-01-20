package mapkha

import (
	"reflect"
	"testing"
)

func TestWithSmallDict(t *testing.T) {
	dict := MakeDict([]string{"กา"})
	wordcut := NewWordcut(dict)
	tokens := wordcut.Segment("กากา")
	expect := []string{"กา", "กา"}
	if !reflect.DeepEqual(expect, tokens) {
		t.Errorf("Expect %q got %q", expect, tokens)
	}
}

func TestSpace(t *testing.T) {
	dict := MakeDict([]string{"กา"})
	wordcut := NewWordcut(dict)
	tokens := wordcut.Segment("ขา ขา")
	expect := []string{"ขา", " ", "ขา"}
	if !reflect.DeepEqual(expect, tokens) {
		t.Errorf("Expect %q got %q", expect, tokens)
	}
}

func TestLatin(t *testing.T) {
	dict := MakeDict([]string{"กา"})
	wordcut := NewWordcut(dict)
	tokens := wordcut.Segment("ขาACขา")
	expect := []string{"ขา", "AC", "ขา"}
	if !reflect.DeepEqual(expect, tokens) {
		t.Errorf("Expect %q got %q", expect, tokens)
	}
}

func TestLoadDefaultDict(t *testing.T) {
	_, err := LoadDefaultDict()
	if err != nil {
		t.Fail()
	}
}

func TestBasic(t *testing.T) {
	dict, _ := LoadDefaultDict()
	wordcut := NewWordcut(dict)
	wlst := wordcut.Segment("กากา")
	expect := []string{"กา", "กา"}
	if !reflect.DeepEqual(expect, wlst) {
		t.Errorf("Expect %q got %q", expect, wlst)
	}
}

func TestBasicUnk(t *testing.T) {
	dict, _ := LoadDefaultDict()
	wordcut := NewWordcut(dict)
	wlst := wordcut.Segment("จะเว")
	expect := []string{"จะ", "เว"}
	if !reflect.DeepEqual(expect, wlst) {
		t.Errorf("Expect %q got %q", expect, wlst)
	}
}

func TestBokWa(t *testing.T) {
	dict, _ := LoadDefaultDict()
	wordcut := NewWordcut(dict)
	wlst := wordcut.Segment("บอกว่า")
	expect := []string{"บอก", "ว่า"}
	if !reflect.DeepEqual(expect, wlst) {
		t.Errorf("Expect %q got %q", expect, wlst)
	}
}

func TestFromLaw(t *testing.T) {
	dict, _ := LoadDefaultDict()
	wordcut := NewWordcut(dict)
	wlst := wordcut.Segment("มาตรา")
	expect := []string{"มาตรา"}

	if !reflect.DeepEqual(expect, wlst) {
		t.Errorf("Expect %q got %q", expect, wlst)
	}
}

func TestBasicDictEdgeBuider(t *testing.T) {
	dict, _ := LoadDefaultDict()
	builder := NewDictEdgeBuilder(dict)
	text := []rune("มาตรา")
	path := make([]*Edge, len(text)+1)

	path[0] = &Edge{S: 0, EdgeType: INIT, WordCount: 0, UnkCount: 0}
	context := EdgeBuildingContext{runes: text, I: 0, Ch: 'ม',
		Path: path, LeftBoundary: 0, BestEdge: nil}
	builder.Build(&context)

	context = EdgeBuildingContext{runes: text, I: 1, Ch: 'า',
		Path: path, LeftBoundary: 0, BestEdge: nil}
	builder.Build(&context)

	path[2] = &Edge{S: 0, EdgeType: DICT, WordCount: 1, UnkCount: 0}
	path[3] = &Edge{S: 0, EdgeType: DICT, WordCount: 1, UnkCount: 0}

	context = EdgeBuildingContext{runes: text, I: 2, Ch: 'ต',
		Path: path, LeftBoundary: 0, BestEdge: nil}
	builder.Build(&context)

	context = EdgeBuildingContext{runes: text, I: 3, Ch: 'ร',
		Path: path, LeftBoundary: 0, BestEdge: nil}
	builder.Build(&context)

	context = EdgeBuildingContext{runes: text, I: 4, Ch: 'า',
		Path: path, LeftBoundary: 0, BestEdge: nil}
	edge := builder.Build(&context)

	if edge == nil {
		t.Errorf("Expect edge not to be nil")
	}

	if edge.S != 0 {
		t.Errorf("Expect edge.S == 0")
	}
}
