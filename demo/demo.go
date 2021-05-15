package main

import (
	"fmt"

	"github.com/kvendrik/stringify"
)

func main() {
	fmt.Println(stringify.Interface(map[string]interface{}{
		"name": "John Doe",
		"age":  26,
		"data": map[string]interface{}{
			"city":  "New York",
			"state": "NY",
		},
		"list": []interface{}{23, "NY", 25},
	}))

	fmt.Println(stringify.Array([]interface{}{23, "NY", 25}))
}
