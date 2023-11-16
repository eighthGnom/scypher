package cypher

type QueryPattern struct {
	Nodes NodePattern
	Edges EdgePattern
}

type NodePattern struct {
	Node *Node
}

type EdgePattern struct {
	Edge *Edge
}

type Condition string

type OrderByOperator string

type Path string

type BooleanOperator string

type Distinct string

type Asterisk string

type AliasOperator string

const (
	// And symbol condition "&"
	AndSymbol Condition = "&"
	// Or symbol condition "|"
	OrSymbol Condition = "|"
	// ":" symbol condition
	Colon Condition = ":"
	Empty Condition = ""

	AND Condition = "AND"
	OR  Condition = "OR"

	LessThan             BooleanOperator = "<"
	GreaterThan          BooleanOperator = ">"
	LessThanOrEqualTo    BooleanOperator = "<="
	GreaterThanOrEqualTo BooleanOperator = ">="
	EqualTo              BooleanOperator = "="
	IN                   BooleanOperator = "IN"
	IS                   BooleanOperator = "IS"
	StartsWith           BooleanOperator = "STARTS WITH"
	EndsWith             BooleanOperator = "ENDS WITH"
	Contains             BooleanOperator = "CONTAINS"

	Asc  OrderByOperator = "ASC"
	Desc OrderByOperator = "DESC"

	// Plain --
	Plain Path = "--"
	// Outgoing -->
	Outgoing Path = "-->"
	// Incoming <--
	Incoming Path = "<--"
	// Bidirectional <-->
	Bidirectional Path = "<-->"

	DISTINCT Distinct      = "DISTINCT"
	WildCard Asterisk      = "*"
	AS       AliasOperator = "AS"
)
