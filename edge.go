package mapkha

// Edge - edge of word graph
type Edge struct {
	S         int
	EdgeType  Etype
	WordCount int
	UnkCount  int
}

// IsBetterThan - comparing this edge to another edge
func (edge *Edge) IsBetterThan(another *Edge) bool {
	if edge == nil {
		return false
	}

	if another == nil {
		return true
	}

	if (edge.UnkCount < another.UnkCount) || ((edge.UnkCount == another.UnkCount) && (edge.WordCount < another.WordCount)) {
		return true
	}

	return false
}

// Improved as Roger Peppe suggested in his tweet
// https://twitter.com/rogpeppe/status/574911374645682176
func GraphToRanges(path []Edge) []TextRange {
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
