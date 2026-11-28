package executor

import "nivrasql/parser"

func (e *Executor) ExecuteDropTable(stmt *parser.DropTableStatement) (*Result, error) {
	err := e.Database.DropTable(stmt.Table)

	if err != nil {
		return nil, err
	}

	return &Result{
		Message: "Table deleted",
	}, nil
}
