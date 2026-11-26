package engine

import "fmt"

func Compare(a any, operator string, b any) (bool, error) {
	switch operator {
	case "==", "=":
		return a == b, nil
	case "!=", "<>":
		return a != b, nil
	case ">":
		switch x := a.(type) {
		case int:
			if y, ok := b.(int); ok {
				return x > y, nil
			}
		case string:
			if y, ok := b.(string); ok {
				return x > y, nil
			}
		case float64:
			if y, ok := b.(float64); ok {
				return x > y, nil
			}
		case bool:
			return false, fmt.Errorf("cannot compare BOOL using >")
		}
	case "<":
		switch x := a.(type) {
		case int:
			if y, ok := b.(int); ok {
				return x < y, nil
			}
		case string:
			if y, ok := b.(string); ok {
				return x < y, nil
			}
		case float64:
			if y, ok := b.(float64); ok {
				return x < y, nil
			}
		case bool:
			return false, fmt.Errorf("cannot compare BOOL using <")
		}
	case ">=":
		switch x := a.(type) {
		case int:
			if y, ok := b.(int); ok {
				return x >= y, nil
			}
		case string:
			if y, ok := b.(string); ok {
				return x >= y, nil
			}
		case float64:
			if y, ok := b.(float64); ok {
				return x >= y, nil
			}
		case bool:
			return false, fmt.Errorf("cannot compare BOOL using >=")
		}
	case "<=":
		switch x := a.(type) {
		case int:
			if y, ok := b.(int); ok {
				return x <= y, nil
			}
		case string:
			if y, ok := b.(string); ok {
				return x <= y, nil
			}
		case float64:
			if y, ok := b.(float64); ok {
				return x <= y, nil
			}
		case bool:
			return false, fmt.Errorf("cannot compare BOOL using <=")
		}
	}
	return false, fmt.Errorf("cannot compare %T and %T using %s", a, b, operator)
}
