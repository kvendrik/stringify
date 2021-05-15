package stringify

func Stringify(data map[string]interface{}) string {
  dataLength := len(data)
  iteration := 0
  currentJSON := "{"
  for key, value := range data {
    iteration += 1
    currentJSON += "\""+key+"\"" + ":" + stringifyValue(value)
    if iteration != dataLength {
      currentJSON += ","
    }
  }
  currentJSON += "}"
  return currentJSON
}

func stringifyValue(value interface{}) string {
  switch value.(type) {
    case string:
      return "\"" + value.(string) + "\""
    case int, float32, float64:
      return "null"
  }

  return Stringify(value.(map[string]interface{}))
}
