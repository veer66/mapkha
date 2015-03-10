package mapkha

type DictAcceptor struct {
	l int
	r int
	final bool
	offset int
	valid bool
	buf []rune
}

const BUFSIZE = 2

func NewDictAcceptor(l int, r int) *DictAcceptor {
	return &DictAcceptor{l, r, false, 0, true, make([]rune, 0, BUFSIZE)}
}

func (a *DictAcceptor) Transit(ch rune, dict [][]rune, idx *DictIndex) {
	var found bool
	a.l, found = DictSeek(LEFT, idx, dict, a.l, a.r, a.offset, ch)
	if found {
		a.r, _ = DictSeek(RIGHT, idx, dict, a.l, a.r, a.offset, ch)
		a.offset++
		w := dict[a.l]
		wlen := len(w)
		a.final = (wlen == a.offset)
		if a.offset <= BUFSIZE {
			a.buf = append(a.buf, ch)
		}
	} else {
		a.valid = false
	}
}
