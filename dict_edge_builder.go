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
	return &DictEdgeBuilder{dict, make([]*dictBuilderPointer, 0, 20)}
}

func (builder *DictEdgeBuilder) updatePointer(pointer *dictBuilderPointer, ch rune) *dictBuilderPointer {
	childNode, found := builder.dict.Lookup(pointer.NodeID, pointer.Offset, ch)
	if !found {
		return nil
	}
	pointer.NodeID = childNode.ChildID
	pointer.Offset += 1
	pointer.IsFinal = childNode.IsFinal
	return pointer
}

// Build - build new edge from dictionary
func (builder *DictEdgeBuilder) Build(context *EdgeBuildingContext) *Edge {
	builder.pointers = append(builder.pointers, &dictBuilderPointer{S: context.I})

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
			s := 1 + context.I - pointer.Offset
			source := context.Path[s]
			edge := &Edge{
				S:         s,
				EdgeType:  DICT,
				WordCount: source.WordCount + 1,
				UnkCount:  source.UnkCount}
			if !bestEdge.IsBetterThan(edge) {
				bestEdge = edge
			}
		}
	}

	return bestEdge
}

func (builder *DictEdgeBuilder) Reset() {
	builder.pointers = builder.pointers[:0]
}
