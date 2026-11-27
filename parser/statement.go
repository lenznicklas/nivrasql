package parser

type Statement interface {
	statementNode()
}

type SelectStatement struct {
	Table   string
	Columns []string
	Where   *Condition
}

type InsertStatement struct {
	Table  string
	Values []any
}

type DeleteStatement struct {
	Table string
	Where *Condition
}

type Condition struct {
	Column   string
	Operator string
	Value    any
}

func (*SelectStatement) statementNode() {}
func (*InsertStatement) statementNode() {}
func (*DeleteStatement) statementNode(){}
