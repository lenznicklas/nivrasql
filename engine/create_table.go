package engine

import "fmt"

func (db *Database) CreateTable(name string, columns []Column) error {
	if _, exists := db.Tables[name]; exists {
		return fmt.Errorf("table %s already exists", name)
	}
	db.Tables[name] = &Table{
		Name:    name,
		Columns: columns,
		Rows:    []Row{},
	}

	return nil
}
