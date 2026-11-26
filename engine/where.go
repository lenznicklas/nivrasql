package engine

import "fmt"

func (t *Table) GetColumnIndex(name string) int {
	for i, col := range t.Columns {
		if col.Name == name {
			return i
		}
	}

	return -1
}

func (t *Table) Where(condition Condition) ([]Row, error) {
	index := t.GetColumnIndex(condition.Column)

	if index == -1 {
		return nil, fmt.Errorf("column %s not found", condition.Column)
	}

	results := []Row{}

	for _, row := range t.Rows {
		comp, err := Compare(row[index], condition.Operator, value)
		if err != nil{
			return nil, err
		}
		if comp {
			results = append(results, row)
		}
	}

	return results, nil
}
