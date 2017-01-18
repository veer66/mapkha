package mapkha

type edgeBuilderFactory func() EdgeBuilder

type Wordcut struct {
	dict                 *Dict
	edgeBuilderFactories []edgeBuilderFactory
}

func NewWordcut(dict *Dict) *Wordcut {
	factories := []edgeBuilderFactory{
		func() EdgeBuilder {
			return NewDictEdgeBuilder(dict)
		},
		func() EdgeBuilder {
			return &UnkEdgeBuilder{}
		}}
	return &Wordcut{dict, factories}
}

func (w *Wordcut) Segment(text string) []string {
	textRunes := []rune(text)
	edgeBuilders := make([]EdgeBuilder, 0)
	for _, factory := range w.edgeBuilderFactories {
		edgeBuilders = append(edgeBuilders, factory())
	}
	path := buildPath(textRunes, edgeBuilders)
	ranges := pathToRanges(path)
	tokens := make([]string, len(ranges))
	for i, r := range ranges {
		tokens[i] = string(textRunes[r.s:r.e])
	}
	return tokens
}
