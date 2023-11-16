package cypher

type PatterBuilder struct {
	patterns []QueryPattern
}

func NewPatternBuilder() *PatterBuilder {
	return &PatterBuilder{}
}

func (pb *PatterBuilder) AddNode(node *Node) *PatterBuilder {
	pb.patterns = append(pb.patterns, node.AsPattern())
	return pb
}

func (pb *PatterBuilder) AddEdge(edge *Edge) *PatterBuilder {
	pb.patterns = append(pb.patterns, edge.AsPattern())
	return pb
}

func (pb *PatterBuilder) ReleasePatterns() []QueryPattern {
	return pb.patterns
}

func (pb *PatterBuilder) Clear() *PatterBuilder {
	pb.patterns = make([]QueryPattern, 0)
	return pb
}
