package mapkha

import (
	"fmt"
)

func buildPath(textRunes []rune, edgeBuilders []EdgeBuilder) []Edge {
	path := make([]Edge, len(textRunes)+1)
	path[0] = Edge{S: 0, EdgeType: INIT, WordCount: 0, UnkCount: 0}
	leftBoundary := 0
	for i, ch := range textRunes {
		fmt.Printf("I=%#v ch=%#v\n", i, ch)
		var bestEdge *Edge
		for _, edgeBuilder := range edgeBuilders {
			context := EdgeBuildingContext{
				runes:        textRunes,
				Path:         path,
				I:            i,
				Ch:           ch,
				LeftBoundary: leftBoundary,
				BestEdge:     bestEdge}
			edge := edgeBuilder.Build(&context)

			if edge.IsBetterThan(bestEdge) {
				bestEdge = edge
			}
		}

		if bestEdge == nil {
			panic("bestEdge must not be nil")
		}

		if bestEdge.EdgeType != UNK {
			leftBoundary = i + 1
		}
	}
	return path
}
