package mapkha

import (
	"reflect"
	"testing"
)

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
