package mapkha

type TextRange struct {
	s int
	e int
}

// Improved as Roger Peppe suggested in his tweet
// https://twitter.com/rogpeppe/status/574911374645682176
func pathToRanges(path []Edge) []TextRange {
	ranges := make([]TextRange, len(path))
	j := len(ranges) - 1
	for e := len(path) - 1; e > 0; {
		s := path[e].S
		ranges[j] = TextRange{s, e}
		j--
		e = s
	}
	return ranges[j+1:]
}
