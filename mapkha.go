package main

import ("fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadDict(path string) ([][]rune, error) {
	b_slice, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := string(b_slice)
	swords := strings.Split(data, "\n")
	rwords := make([][]rune, len(swords))
	for i, word := range swords {
		rwords[i] = []rune(word)
	}
	return rwords, nil
}

type DictAcceptor struct {
	l int
	r int
	final bool
	offset int
	valid bool
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

func DictSeek(policy int, dict [][]rune, l int, r int, offset int, ch rune) (int, bool) {
	ans := 0
	found := false
	m := 0	
	if policy != LEFT && policy != RIGHT {
		return 0, found
	}	
	for ;l<=r; {
		m = (l+r) / 2
		w := dict[m]
		wlen := len(w)
		if wlen <= offset {
			l = m + 1
		} else {
			ch_ := w[offset]
			if ch_ < ch {
				l = m + 1
			} else if ch_ > ch {
				r = m - 1
			} else {
				ans = m
				found = true
				switch policy {
				case LEFT: r = m - 1
				case RIGHT: l = m + 1
				}
			}			
		}
	}	
	return ans, found
}

type TextRange struct {
	s int
	e int
}

type Edge struct {
	w int
	unk int
	p int
}

const (
	LEFT = 1
	RIGHT = 2
)

func TransitAll(acc []DictAcceptor, ch rune, dict [][]rune) []DictAcceptor {
	_acc := append(acc, DictAcceptor{0, len(dict)-1, false, 0, true})
	__acc := make([]DictAcceptor, 0, len(_acc))
	for _, a := range(_acc) {
		a.Transit(ch, dict)
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

func BuildEdges(i int, acc []DictAcceptor, g []Edge) []Edge {
	edges := make([]Edge, 0, len(acc))
	for _, a := range(acc) {
		if a.final {
			p := i - a.offset + 1
			src := g[p]
			edge := Edge{src.w + 1, src.unk, p}
			edges = append(edges, edge)			
		}
	}

	if len(edges) == 0 {
		edge := Edge{100, 100, 0}
		edges = append(edges, edge)
	}
	return edges
}

func BuildGraph(t []rune, dict [][]rune) []Edge {
	g := make([]Edge, len(t) + 1)
	g[0] = Edge{0, 0, -1}
	var acc []DictAcceptor
	for i, ch := range(t) {
		acc = TransitAll(acc, ch, dict)
		edges := BuildEdges(i, acc, g)
		e := BestEdge(edges)
		g[i+1] = *e 
	}
	return g
}

// Improved as Roger Peppe suggested in his tweet
// https://twitter.com/rogpeppe/status/574911374645682176
func GraphToRanges(g []Edge) []TextRange {
	ranges := make([]TextRange, len(g))
	j := len(ranges)-1
	for e := len(g) - 1; e > 0; {
		s := g[e].p
		ranges[j] = TextRange{s, e}
		j--
		e = s
	}
	return ranges[j:]
}

func Segment(_t string, dict [][]rune) []string {
	t := []rune(_t)
	ranges := GraphToRanges(BuildGraph(t, dict)) 
	wlst := make([]string, len(ranges))
	for i, r := range ranges {
		wlst[i] = string(t[r.s:r.e])
	}
	return wlst
}

func main() {
	dict, e := LoadDict("tdict-std.txt")
	check(e)
	wl := Segment("ตัดคำไหม", dict)
	fmt.Println(wl)
}
