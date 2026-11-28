package executor

import (
	"fmt"
	"nivrasql/engine"
	"nivrasql/parser"
)

func (e *Executor) Execute(stmt parser.Statement) (*Result, error) {
	switch s := stmt.(type) {
	case *parser.CreateTableStatement:
		return e.ExecuteCreateTable(s)
	case *parser.InsertStatement:
		return e.ExecuteInsert(s)
	case *parser.DropTableStatement:
		return e.ExecuteDropTable(s)
	case *parser.SelectStatement:
		return e.ExecuteSelect(s)
	default:
		return nil, fmt.Errorf("unsupported statement: %T", stmt)
	}
}

func (e *Executor) ExecuteSelect(stmt *parser.SelectStatement) (*Result, error) {

	var condition *engine.Condition

	if stmt.Where != nil {
		condition = &engine.Condition{
			Column:   stmt.Where.Column,
			Operator: stmt.Where.Operator,
			Value:    stmt.Where.Value,
		}
	}

	rows, err := e.Database.Select(stmt.Columns, condition)
	if err != nil {
		return nil, err
	}

	cols := stmt.Columns

	if len(cols) == 1 && cols[0] == "*" {
		cols := []string{}

		for _, col := range table.Columns {
			cols = append(cols, col.Name)
		}
	}

	return &Result{
		Columns: cols,
		Rows:    rows,
	}, nil
}
