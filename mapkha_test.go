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
	wlst := Segment("กากา", dict, nil)
	if wlst[0] != "กา" || wlst[0] != "กา" || len(wlst) != 2 {
		t.Fail()
	}
}

func TestIdx(t *testing.T) {
	dict, _ := LoadDefaultDict()
	idx := MakeIndex(dict)
	wlst := Segment("กากา", dict, idx)
	if wlst[0] != "กา" || wlst[0] != "กา" || len(wlst) != 2 {
		t.Fail()
	}
}


func TestBasicUnk(t *testing.T) {
	dict, _ := LoadDefaultDict()
	wlst := Segment("จะเว", dict, nil)
	if wlst[0] != "จะ" || wlst[1] != "เว" || len(wlst) != 2 {
		t.Fail()
	}
}


