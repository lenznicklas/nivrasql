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
	t.Rows = append(t.Rows, row)
	return nil
}

func ValidateType(value any, dataType string) bool {
	switch dataType {
	case "INT":
		_, ok := value.(int)
		return ok

	case "TEXT":
		_, ok := value.(string)
		return ok

	case "BOOL":
		_, ok := value.(bool)
		return ok

	case "FLOAT":
		_, ok := value.(float64)
		return ok
	}

	return false
}
