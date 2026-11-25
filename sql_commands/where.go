package sqlcommands

import "fmt"

func (t *Table) GetColumnIndex(name string) int {
	for i, col := range t.Columns {
		if col.Name == name {
			return i
		}
	}

	return -1
}

func (t *Table) Where(column string, value any) ([]Row, error) {
	index := t.GetColumnIndex(column)

	if index == -1 {
		return nil, fmt.Errorf("column %s not found", column)
	}

	results := []Row{}

	for _, row := range t.Rows {
		if row[index] == value {
			results = append(results, row)
		}
	}

	return results, nil
}
