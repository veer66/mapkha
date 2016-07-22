package mapkha

import (
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
	if wlst[0] != "กา" || wlst[0] != "กา" || len(wlst) != 2 {
		t.Fail()
	}
}

func TestBasicUnk(t *testing.T) {
	dict, _ := LoadDefaultDict()
	wordcut := NewWordcut(dict)
	wlst := wordcut.Segment("จะเว")
	if wlst[0] != "จะ" || wlst[1] != "เว" || len(wlst) != 2 {
		t.Fail()
	}
}
