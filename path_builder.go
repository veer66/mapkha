package mapkha

var globalContext = &EdgeBuildingContext{}

func buildPath(textRunes []rune, edgeBuilders []EdgeBuilder) []*Edge {
	path := make([]*Edge, len(textRunes)+1)
	path[0] = &Edge{S: 0, EdgeType: INIT, WordCount: 0, UnkCount: 0}
	leftBoundary := 0
	for i, ch := range textRunes {
		var bestEdge *Edge
		for _, edgeBuilder := range edgeBuilders {
			globalContext.runes = textRunes
			globalContext.Path = path
			globalContext.I = i
			globalContext.Ch = ch
			globalContext.LeftBoundary = leftBoundary
			globalContext.BestEdge = bestEdge

			edge := edgeBuilder.Build(globalContext)

			if edge != nil && edge.IsBetterThan(bestEdge) {
				bestEdge = edge
			}
		}

		if bestEdge == nil {
			panic("bestEdge must not be nil")
		}

		if bestEdge.EdgeType != UNK {
			leftBoundary = i + 1
		}

		path[i+1] = bestEdge
	}
	return path
}
