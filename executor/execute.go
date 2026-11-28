package executor

import (
	"fmt"
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

}
