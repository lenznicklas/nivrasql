package sqlcommands

import "fmt"

func (db *Database) Insert(name string, row Row) error {
	table, exists := db.Tables[name]

	if !exists {
		return fmt.Errorf("table %s does not exist", name)
	}

	return table.Insert(row)
}

func (t *Table) Insert(row Row) error {
	t.Rows = append(t.Rows, row)
	return nil
}
