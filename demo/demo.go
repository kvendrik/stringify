package main

import (
	"fmt"

	"github.com/kvendrik/stringify"
)

func main() {
	fmt.Println(stringify.Stringify(map[string]interface{}{
		"name": "John Doe",
		"age":  26,
		"data": map[string]interface{}{
			"city":  "New York",
			"state": "NY",
		},
	}))
}
