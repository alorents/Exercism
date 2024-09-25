package flatten

func Flatten(nested interface{}) []interface{} {
	result := make([]interface{}, 0)

	switch i := nested.(type) {
	case []interface{}:
		for _, val := range i {
			result = append(result, Flatten(val)...)
		}

	case interface{}:
		result = append(result, i)
	}

	return result
}
