package mapkha

type l0idx map[rune]int
type l1idx map[rune](l0idx)

type Index struct {
	left0 l0idx
	right0 l0idx
	left1 l1idx
	right1 l1idx
}

func MakeIndex(d *Dict) *Index {
	l := make(l0idx)
	r := make(l0idx)

	l1 := make(l1idx)
	r1 := make(l1idx)

	for i, w := range d.GetSlice() {
		if len(w) > 0 {
			_, exist := l[w[0]]
			if !exist {
				l[w[0]] = i
			}
			r[w[0]] = i

			if len(w) > 1 {
				_, e_l1_0 := l1[w[0]]
				if !e_l1_0 {
					l1[w[0]] = make(l0idx)
				}
				_, e_r1_0 := r1[w[0]]
				if !e_r1_0 {
					r1[w[0]] = make(l0idx)
				}

				l1_1, _ := l1[w[0]]
				_, exist := l1_1[w[1]]
				if !exist {
					l1_1[w[1]] = i
				}
				
				r1_1, _ := r1[w[0]]
				r1_1[w[1]] = i
			}
		}
	}
	return &Index{l, r, l1, r1}
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

func (idx *Index) Get1(policy int, ch0 rune, ch1 rune) (int, bool) {
	var i int
	var found bool
	var x0 *l1idx
	switch policy {
	case LEFT:
		x0 = &idx.left1
		
	case RIGHT:
		x0 = &idx.right1		
	}
	x1 := (*x0)[ch0]
	i, found = x1[ch1]
	return i, found
}

