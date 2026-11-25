package engine

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
