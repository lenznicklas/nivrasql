package executor

import (
	"nivrasql/engine"
	"nivrasql/parser"
)

func (e *Executor) ExecuteCreateTable(stmt *parser.CreateTableStatement) (*Result, error) {
	cols := []engine.Column{}

	for _, col := range stmt.Columns {
		cols = append(cols, engine.Column{
			Name: col.Name,
			Type: col.Type,
		})
	}

	err := e.Database.CreateTable(
		stmt.Table,
		cols,
	)

	if err != nil {
		return nil, err
	}

	return &Result{
		Message: "Table created",
	}, nil
}
