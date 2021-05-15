package stringify

func Interface(data map[string]interface{}) string {
	dataLength := len(data)
	iteration := 0
	currentJSON := "{"
	for key, value := range data {
		iteration += 1
		currentJSON += "\"" + key + "\"" + ":" + stringifyValue(value)
		if iteration != dataLength {
			currentJSON += ","
		}
	}
	currentJSON += "}"
	return currentJSON
}

func Array(data []interface{}) string {
	dataLength := len(data)
	iteration := 0
	currentJSON := "["
	for _, value := range data {
		iteration += 1
		currentJSON += stringifyValue(value)
		if iteration != dataLength {
			currentJSON += ","
		}
	}
	currentJSON += "]"
	return currentJSON
}

func stringifyValue(value interface{}) string {
	switch value.(type) {
	case string:
		return "\"" + value.(string) + "\""
	case []interface{}:
		v, _ := value.([]interface{})
		return Array(v)
	case map[string]interface{}:
		return Interface(value.(map[string]interface{}))
	case int, float32, float64:
		v, _ := value.(int)
		return integerToString(v)
	}

	return "null"
}

func integerToString(i int) string {
	bytes := []byte{}
	cur := i
	for {
		digit := cur % 10
		cur = cur / 10
		bytes = append([]byte{byte(digit + 48)}, bytes...)
		if cur <= 0 {
			break
		}
	}

	return string(bytes)
}
