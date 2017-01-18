package mapkha

type DictAcceptor struct {
	p      int
	offset int
	final  bool
	valid  bool
}

// Reset - reset internal state
func (a *DictAcceptor) Reset(p int) {
	a.p = p
	a.final = false
	a.offset = 0
	a.valid = true
}

// Transit - walk on prefix tree by new rune
func (a *DictAcceptor) Transit(ch rune, dict *Dict) {
	pointer, found := dict.Lookup(a.p, a.offset, ch)

	if !found {
		a.valid = false
		return
	}

	a.p = pointer.ChildID
	a.offset++
	a.final = pointer.IsFinal
}

// AccPool - pool of dict acceptor
type AccPool struct {
	acc []DictAcceptor
	i   int
}

// NewAccPool - build acceptor pool
func NewAccPool() *AccPool {
	return &AccPool{make([]DictAcceptor, 0, 4096), 0}
}

// Reset - reset acceptor pool
func (pool *AccPool) Reset() {
	pool.i = 0
}

// Obtain - obtain dict acceptor at p
func (pool *AccPool) Obtain(p int) *DictAcceptor {
	if pool.i >= len(pool.acc) {
		pool.acc = append(pool.acc, DictAcceptor{})
	}

	a := &pool.acc[pool.i]
	a.Reset(p)
	pool.i++
	return a
}
