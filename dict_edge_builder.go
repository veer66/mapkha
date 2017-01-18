package mapkha

import "fmt"

type DictEdgeBuilder struct {
	dict     *Dict
	pointers []*dictBuilderPointer
}

type dictBuilderPointer struct {
	NodeID  int
	S       int
	Offset  int
	IsFinal bool
}

func NewDictEdgeBuilder(dict *Dict) *DictEdgeBuilder {
	return &DictEdgeBuilder{dict, make([]*dictBuilderPointer, 0)}
}

func (builder *DictEdgeBuilder) updatePointer(pointer *dictBuilderPointer, ch rune) *dictBuilderPointer {
	childNode, found := builder.dict.Lookup(pointer.NodeID, pointer.Offset, ch)
	if !found {
		return nil
	}
	return &dictBuilderPointer{NodeID: pointer.NodeID,
		S:       pointer.S,
		Offset:  pointer.Offset + 1,
		IsFinal: childNode.IsFinal}
}

// Build - build new edge from dictionary
func (builder *DictEdgeBuilder) Build(context *EdgeBuildingContext) *Edge {
	oldPointers := append(builder.pointers,
		&dictBuilderPointer{
			NodeID:  0,
			S:       context.I,
			Offset:  0,
			IsFinal: false})

	fmt.Printf("@@@ %#v\n", oldPointers)
	pointers := make([]*dictBuilderPointer, 0)

	// (->> (map updatePointer) (remove nil))
	for _, pointer := range oldPointers {
		newPointer := builder.updatePointer(pointer, context.Ch)
		if newPointer != nil {
			pointers = append(pointers, newPointer)
		}
	}

	builder.pointers = pointers
	var bestEdge *Edge

	for _, pointer := range pointers {
		fmt.Printf("P %#v\n", pointer)
		if pointer.IsFinal {
			source := context.Path[pointer.S]
			edge := &Edge{
				S:         pointer.S,
				EdgeType:  DICT,
				WordCount: source.WordCount + 1,
				UnkCount:  source.UnkCount}
			fmt.Printf("DICT-EDGE %#v\n", edge)
			if edge.IsBetterThan(bestEdge) {
				bestEdge = edge
			}
		}
	}

	return bestEdge
}
