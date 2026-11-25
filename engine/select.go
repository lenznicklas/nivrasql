package engine

import (
	"fmt"
)

func (t *Table) Select(columns []string) ([]Row, error) {
	if len(columns) == 1 && columns[0] == "*" {
		return t.Rows, nil
	}
	indices := []int{}
	result := []Row{}

	for _, column := range columns {
		index := t.GetColumnIndex(column)
		if index == -1 {
			return nil, fmt.Errorf("column %s not found", column)
		}

		indices = append(indices, index)
	}

	for _, row := range t.Rows {
		newRow := Row{}

		for _, index := range indices {
			newRow = append(newRow, row[index])
		}

		result = append(result, newRow)
	}

	return result, nil
}
