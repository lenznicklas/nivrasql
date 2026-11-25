package engine

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
