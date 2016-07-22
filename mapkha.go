package mapkha

type TextRange struct {
	s int
	e int
}

type Edge struct {
	w     int
	unk   int
	p     int
	etype int
}

type Wordcut struct {
	dict *Dict
	pool *AccPool
}

func NewWordcut(dict *Dict) *Wordcut {
	return &Wordcut{dict, NewAccPool()}
}

func (w *Wordcut) TransitAll(acc []*DictAcceptor, ch rune) []*DictAcceptor {
	_acc := append(acc, w.pool.Obtain(0, w.dict.R()))
	__acc := make([]*DictAcceptor, 0, len(_acc))
	for _, a := range _acc {
		a.Transit(ch, w.dict)
		if a.valid {
			__acc = append(__acc, a)
		}
	}
	return __acc
}

func Better(a *Edge, b *Edge) bool {
	if a.unk < a.unk || a.w < b.w {
		return true
	}
	return false
}

func BestEdge(edges []Edge) *Edge {
	l := len(edges)
	if l == 0 {
		return nil
	}
	e := &edges[0]
	for i := 1; i < l; i++ {
		if Better(&edges[i], e) {
			e = &edges[i]
		}
	}
	return e
}

func BuildEdges(i int, acc []*DictAcceptor, g []Edge, left int) []Edge {
	edges := make([]Edge, 0, len(acc))
	for _, a := range acc {
		if a.final {
			p := i - a.offset + 1
			src := g[p]
			edge := Edge{src.w + 1, src.unk, p, DICT}
			edges = append(edges, edge)
		}
	}

	if len(edges) == 0 {
		src := g[left]
		edge := Edge{src.w + 1, src.unk + 1, left, UNK}
		edges = append(edges, edge)
	}
	return edges
}

func (w *Wordcut) BuildGraph(t []rune) []Edge {
	g := make([]Edge, len(t)+1)
	g[0] = Edge{0, 0, -1, INIT}
	var acc []*DictAcceptor
	left := 0
	for i, ch := range t {
		acc = w.TransitAll(acc, ch)
		edges := BuildEdges(i, acc, g, left)
		e := BestEdge(edges)
		if e.etype != UNK {
			left = i + 1
		}
		g[i+1] = *e
	}
	return g
}

// Improved as Roger Peppe suggested in his tweet
// https://twitter.com/rogpeppe/status/574911374645682176
func GraphToRanges(g []Edge) []TextRange {
	ranges := make([]TextRange, len(g))
	j := len(ranges) - 1
	for e := len(g) - 1; e > 0; {
		s := g[e].p
		ranges[j] = TextRange{s, e}
		j--
		e = s
	}
	return ranges[j+1:]
}

func (w *Wordcut) Segment(_t string) []string {
	t := []rune(_t)
	ranges := GraphToRanges(w.BuildGraph(t))
	wlst := make([]string, len(ranges))
	for i, r := range ranges {
		wlst[i] = string(t[r.s:r.e])
	}
	w.pool.Reset()
	return wlst
}
