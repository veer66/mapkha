package mapkha

type DictAcceptor struct {
	l int
	r int
	final bool
	offset int
	valid bool
}

func NewDictAcceptor(l int, r int) *DictAcceptor {
	return &DictAcceptor{l, r, false, 0, true}
}

func (a *DictAcceptor) Transit(ch rune, dict [][]rune) {
	var found bool
	a.l, found = DictSeek(LEFT, dict, a.l, a.r, a.offset, ch)
	if found {
		a.r, _ = DictSeek(RIGHT, dict, a.l, a.r, a.offset, ch)
		a.offset++
		w := dict[a.l]
		wlen := len(w)
		a.final = (wlen == a.offset)
	} else {
		a.valid = false
	}
}
