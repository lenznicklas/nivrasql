package engine

import "fmt"

func (t *Table) Delete(condition *Condition) (int, error) {
	if condition == nil {
		count := len(t.Rows)
		t.Rows = []Row{}
		return count, nil
	}

	index := t.GetColumnIndex(condition.Column)

	if index == -1 {
		return 0, fmt.Errorf("column %s not found", condition.Column)
	}

	remaining := []Row{}
	deleted := 0

	for _, row := range t.Rows {
		match, err := Compare(
			row[index],
			condition.Operator,
			condition.Value,
		)

		if err != nil {
			return 0, err
		}

		if match {
			deleted++
		} else {
			remaining = append(remaining, row)
		}
	}

	t.Rows = remaining

	return deleted, nil
}
