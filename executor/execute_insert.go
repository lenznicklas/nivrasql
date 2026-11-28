package executor

import "nivrasql/parser"

func (e *Executor) ExecuteInsert(stmt *parser.InsertStatement) (*Result, error) {
	vals := stmt.Values

	err := e.Database.Insert(stmt.Table, vals)

	if err != nil {
		return nil, err
	}

	return &Result{
		RowsAffected: 1,
		Message:      "Row inserted",
	}, nil
}
