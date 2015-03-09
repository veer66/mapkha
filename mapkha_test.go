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
	wlst := Segment("กากา", dict)
	if wlst[0] != "กา" || wlst[0] != "กา" || len(wlst) != 2 {
		t.Fail()
	}
}

