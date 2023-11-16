package cypher

import (
	"fmt"
)

type Edge struct {
	variable   string
	label      Label
	properties Properties
	path       Path
	condition  Condition
}

func NewEdge() *Edge {
	return &Edge{
		variable:   "",
		label:      Label{},
		properties: Properties{},
		path:       "",
	}
}

func (e *Edge) SetVariable(variable string) *Edge {
	e.variable = variable
	return e
}

func (e *Edge) SetLabel(label string) *Edge {
	e.label.Names = append(e.label.Names, label)
	return e
}

func (e *Edge) SetLabels(condition Condition, labels ...string) *Edge {
	e.label.Names = append(e.label.Names, labels...)
	e.condition = condition
	return e
}

func (e *Edge) SetProps(props ...Property) *Edge {
	for _, p := range props {
		e.properties[p.Key] = p.Value
	}
	return e
}

func (e *Edge) SetPath(path Path) *Edge {
	e.path = path
	return e
}

func (e Edge) AsPattern() QueryPattern {
	return QueryPattern{Edges: EdgePattern{Edge: &e}}
}

func (e Edge) ToCypher() string {
	edge := ""
	if e.variable != "" {
		edge += e.variable
	}
	edge += e.label.ToCypher()
	if len(e.properties) > 0 {
		edge += fmt.Sprintf(" %s", e.properties.ToCypher())
	}
	edge = fmt.Sprintf("-[%v]-", edge)
	switch e.path {
	case Outgoing:
		edge += ">"
	case Incoming:
		edge = "<" + edge
	case Bidirectional:
		edge = "<" + edge + ">"
	case Plain:
	}
	return edge
}
