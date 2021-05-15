# [WIP] Tiny JSON Stringify

## Why?

When building applications that use Go and WebAssembly I often want to be able to send complicated results back to the client.

There's two problems that occur when trying to do this:

1. WebAssembly only supports numbers (int32, int64, float32, and float64). This can be solved by sending back the pointer + length of the string and retreiving it from memory.
2. [`tinygo`](https://tinygo.org/) which is my complier of choice for small WASM binaries doesn't support the `encoding/json` package. Even if it did it would add a fair bit of weight to the binary.

This package aims to solve problem number 2 in the most lightweight way possible (no imported libraries).

## Example

```go
package main

import "fmt"
import "github.com/kvendrik/stringify"

func main() {
  fmt.Println(stringify.Stringify(map[string]interface{}{
    "name": "John Doe",
    "age": 26,
    "data": map[string]interface{}{
      "city": "New York",
      "state": "NY",
    },
  }))
}
```

Yields

```json
{ "name": "John Doe", "age": 26, "data": { "city": "New York", "state": "NY" } }
```

## Progress

- [x] Support strings
- [x] Support interfaces
- [x] Support nested interfaces
- [ ] Support numbers
- [ ] Support arrays
