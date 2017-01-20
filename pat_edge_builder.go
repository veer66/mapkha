package mapkha

type runePatPredicate func(rune) bool

type PatEdgeBuilder struct {
	s        int
	foundS   bool
	foundE   bool
	isPat    runePatPredicate
	edgeType Etype
}

func (builder *PatEdgeBuilder) Build(context *EdgeBuildingContext) *Edge {
	if !builder.foundS {
		if builder.isPat(context.Ch) {
			builder.s = context.I
			builder.foundS = true
		}
	}

	if builder.foundS {
		if builder.isPat(context.Ch) {
			if len(context.runes) == context.I+1 ||
				!builder.isPat(context.runes[context.I+1]) {
				builder.foundE = true
			}
		} else {
			builder.foundS = false
			builder.foundE = false
		}
	}

	if builder.foundS && builder.foundE {
		source := context.Path[builder.s]
		builder.foundS = false
		builder.foundE = false
		return &Edge{S: builder.s,
			EdgeType:  builder.edgeType,
			WordCount: source.WordCount + 1,
			UnkCount:  source.UnkCount}
	}
	return nil
}

func (builder *PatEdgeBuilder) Reset() {
	builder.foundS = false
	builder.foundE = false
	builder.s = 0
}
