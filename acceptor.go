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

func (a *DictAcceptor) Reset(l int, r int) {
	a.l = l
	a.r = r
	a.final = false
	a.offset = 0
	a.valid = true
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

type AccPool struct {
	acc []DictAcceptor
	i int
}

func NewAccPool() *AccPool {
	return &AccPool{make([]DictAcceptor,0, 4096), 0}
}

func (pool *AccPool) Reset() {
	pool.i = 0
}

func (pool *AccPool) Obtain(l int, r int) *DictAcceptor {
	if pool.i < len(pool.acc) {
		a := &pool.acc[pool.i]
		a.Reset(l, r)
		pool.i++
		return a
	} else {
		a := NewDictAcceptor(l, r)
		pool.acc = append(pool.acc, *a)
		pool.i++
		return a
	}
}
