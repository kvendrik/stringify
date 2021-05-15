# [WIP] Tiny JSON Stringify (`7.5kb`)

**Tiny JSON Stringify library for Go. Built to have a way to stringify interfaces in WebAssembly applications (it's `.wasm` file is `7.5kb`).**

## Why?

When building applications that use Go and WebAssembly I often want to be able to send complicated results back to the client.

There's two problems that occur when trying to do this:

1. WebAssembly only supports numbers (int32, int64, float32, and float64). This can be solved by sending back the pointer + length of the string and retreiving it from memory.
2. [`tinygo`](https://tinygo.org/) which is my complier of choice for small WASM binaries doesn't support the `encoding/json` package. Even if it did it would add a fair bit of weight to the binary.

This package aims to solve problem number 2 in the most lightweight way possible (no imported libraries).

## Usage

```
go get github.com/kvendrik/stringify
```

```go
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
```

Yields

```json
{
  "name": "John Doe",
  "age": 26,
  "data": { "city": "New York", "state": "NY" },
  "list": [23, "NY", 25]
}
[23, "NY", 25]
```

## Running the demo

```
make run_demo
```

## Progress

- [x] Support strings
- [x] Support interfaces
- [x] Support nested interfaces
- [x] Support numbers
- [x] Support arrays
- [ ] Support floats
