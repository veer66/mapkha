package mapkha

type l0idx map[rune]int
type l1idx map[rune](l0idx)

type Index struct {
	left0  l0idx
	right0 l0idx
}

func MakeIndex(d *Dict) *Index {
	l := make(l0idx)
	r := make(l0idx)

	for i, w := range d.GetSlice() {
		if len(w) > 0 {
			_, exist := l[w[0]]
			if !exist {
				l[w[0]] = i
			}
			r[w[0]] = i
		}
	}
	return &Index{l, r}
}

func (idx *Index) Get0(policy int, ch rune) (int, bool) {
	var i int
	var found bool
	switch policy {
	case LEFT:
		i, found = idx.left0[ch]
	case RIGHT:
		i, found = idx.right0[ch]
	}
	return i, found
}
