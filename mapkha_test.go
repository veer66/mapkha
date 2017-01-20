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
