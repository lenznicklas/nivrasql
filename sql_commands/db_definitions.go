package sqlcommands

type Column struct {
	Name string
	Type string
}

type Row []any

type Table struct {
	Name    string
	Columns []Column
	Rows    []Row
}

type Database struct {
	Tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]*Table),
	}
}
