package executor

import (
	"nivrasql/engine"
	"nivrasql/lexer"
	"nivrasql/parser"
)

type Executor struct {
	Database *engine.Database
}

func NewExecutor(db *engine.Database) *Executor {
	return &Executor{
		Database: db,
	}
}

type Result struct {
	Columns      []string
	Rows         []engine.Row
	RowsAffected int
	Message      string
}

func (e *Executor) ExecuteSQL(command string) (*Result, error) {
	l := lexer.Lexer{
		Command: command,
	}

	if err := l.Lex(); err != nil {
		return nil, err
	}

	p := parser.NewParser(l.Tokens)

	stmt, err := p.Parse()
	if err != nil {
		return nil, err
	}

	return e.Execute(stmt)
}
