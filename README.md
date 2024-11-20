# valdo

A Go package for validating JSON that can generate JSON Schema.

Features:

* Supports the latest JSON Schema specification (2020-12).
* Pure Go.
* No code generation, no reflection, no unsafe code.
* User-freindly error messages.
* Concurrency-safe, no global state.
* Strict by default, without imlicit type casting.

## Installation

```bash
go get github.com/orsinium-labs/valdo
```

## Usage

```go
validator := valdo.Object(
    valdo.Property("name", valdo.String(valdo.MinLen(1))),
    valdo.Property("admin", valdo.Bool()),
)

// validate JSON
input := []byte(`{"name": "aragorn", "admin": true}`)
err := valdo.Validate(validator, raw)

// validate and unmarshal JSON
type User struct {
    Name  string `json:"name"`
    Admin bool   `json:"admin"`
}
err := valdo.Unmarshal[User](validator, input)

// generate JSON Schema
schema := valdo.Schema(validator)
```
