package mapkha

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
	pointer.Offset += 1
	pointer.IsFinal = childNode.IsFinal
	return pointer
}

// Build - build new edge from dictionary
func (builder *DictEdgeBuilder) Build(context *EdgeBuildingContext) *Edge {
	builder.pointers = append(builder.pointers,
		&dictBuilderPointer{
			NodeID:  0,
			S:       context.I,
			Offset:  0,
			IsFinal: false})

	//pointers := make([]*dictBuilderPointer, 0)

	// (->> (map updatePointer) (remove nil))
	newIndex := 0
	for i, _ := range builder.pointers {
		newPointer := builder.updatePointer(builder.pointers[i], context.Ch)
		if newPointer != nil {
			builder.pointers[newIndex] = newPointer
			newIndex++
		}
	}

	builder.pointers = builder.pointers[:newIndex]
	var bestEdge *Edge

	for _, pointer := range builder.pointers {
		if pointer.IsFinal {
			source := context.Path[pointer.S]
			edge := &Edge{
				S:         pointer.S,
				EdgeType:  DICT,
				WordCount: source.WordCount + 1,
				UnkCount:  source.UnkCount}
			if edge.IsBetterThan(bestEdge) {
				bestEdge = edge
			}
		}
	}

	return bestEdge
}
