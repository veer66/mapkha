package mapkha

type EdgeBuildingContext struct {
	runes        []rune
	I            int
	Ch           rune
	Path         []Edge
	LeftBoundary int
	BestEdge     *Edge
}

type EdgeBuilder interface {
	Build(*EdgeBuildingContext) *Edge
}
