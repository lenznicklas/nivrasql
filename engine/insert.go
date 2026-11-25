package engine

import "fmt"

func (db *Database) Insert(name string, row Row) error {
	table, exists := db.Tables[name]

	if !exists {
		return fmt.Errorf("table %s does not exist", name)
	}

	return table.Insert(row)
}

func (t *Table) Insert(row Row) error {
	if len(row) != len(t.Columns) {
		return fmt.Errorf("invalid number of values")
	}
	for n, element := range row {
		if !ValidateType(element, t.Columns[n].Type) {
			return fmt.Errorf("incorrect types")
		}
	}
	t.Rows = append(t.Rows, row)
	return nil
}
