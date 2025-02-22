package utils

import "strconv"

func InterfaceToString(input interface{}) string {
	out := ""

	if input == nil {
		return out
	}

	out = input.(string)

	return out
}

func InterfaceToInt(input interface{}) int {
	out := 0

	if input == nil {
		return out
	}

	switch input.(type) {
	case string:
		data := input.(string)
		out, _ = strconv.Atoi(data)
	}

	return out
}

func InterfaceToFloat64(input interface{}) float64 {
	out := 0.0

	if input == nil {
		return out
	}

	out = input.(float64)

	return out
}

func ArrayinterfaceToArrayString(input []interface{}) []string {
	out := []string{}

	if input == nil {
		return out
	}

	for _, each := range input {
		val := InterfaceToString(each)
		out = append(out, val)
	}

	return out
}

func ArrayInterfaceToString(input []interface{}) string {
	res := ""

	if input == nil {
		return ""
	}

	for i, each := range input {
		val := InterfaceToString(each)
		if i == (len(input) - 1) {
			res += val
		} else {
			res += val + ","
		}

	}

	return res

}
