package engine

import (
	"fmt"
)

func (db *Database) Select(
	name string,
	columns []string,
	condition *Condition,
) ([]Row, error) {
	table, exists := db.Tables[name]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", name)
	}

	return table.Select(columns, condition)
}

func (t *Table) Select(columns []string, condition *Condition) ([]Row, error) {
	rows := t.Rows

	if condition != nil {
		filtered, err := t.Where(*condition)
		if err != nil {
			return nil, err
		}

		rows = filtered
	}

	if len(columns) == 1 && columns[0] == "*" {
		return rows, nil
	}

	indices := []int{}

	for _, column := range columns {
		index := t.GetColumnIndex(column)

		if index == -1 {
			return nil, fmt.Errorf("column %s not found", column)
		}

		indices = append(indices, index)
	}

	result := []Row{}

	for _, row := range t.Rows {
		newRow := Row{}

		for _, index := range indices {
			newRow = append(newRow, row[index])
		}

		result = append(result, newRow)
	}

	return result, nil
}
