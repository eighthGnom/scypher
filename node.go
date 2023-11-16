package cypher

import (
	"fmt"
)

type Node struct {
	variable   string
	label      Label
	properties Properties
}

func NewNode() *Node {
	return &Node{
		label:      Label{},
		properties: Properties{},
		variable:   "",
	}
}

func (n *Node) SetVariable(variable string) *Node {
	n.variable = variable
	return n
}

func (n *Node) SetProps(props ...Property) *Node {
	for _, p := range props {
		n.properties[p.Key] = p.Value
	}
	return n
}

func (n *Node) SetLabel(label string) *Node {
	n.label.Names = append(n.label.Names, label)
	return n
}

func (n *Node) SetLabels(condition Condition, labels ...string) *Node {
	n.label.Names = append(n.label.Names, labels...)
	n.label.Condition = condition
	return n
}

func (n Node) AsPattern() QueryPattern {
	return QueryPattern{Nodes: NodePattern{Node: &n}}
}

func (n Node) ToCypher() (string, error) {
	if n.variable == "" && len(n.label.Names) > 0 {
		return "", fmt.Errorf("node must have a variable with at least one label")
	}
	node := ""
	if n.variable != "" {
		node += n.variable
	}
	node += n.label.ToCypher()
	if len(n.properties) > 0 {
		node += fmt.Sprintf(" %s", n.properties.ToCypher())
	}
	node = fmt.Sprintf("(%s)", node)
	return node, nil
}
