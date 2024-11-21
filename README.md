# valdo

[ [📚 docs](https://pkg.go.dev/github.com/orsinium-labs/valdo/valdo) ] [ [🐙 github](https://github.com/orsinium-labs/valdo) ]

Go package for validating JSON. Can generate [JSON Schema](https://json-schema.org/overview/what-is-jsonschema) (100% compatible with [OpenAPI](https://swagger.io/specification/)), produces user-friendly errors, supports translations.

You could write OpenAPI documentation by hand (which is very painfull) and then use it to validate user input in your HTTP service, but then error messages are very confusing, not user-friendly, and only in English. Or you could write input validaion by hand and then maintain the OpenAPI documentation separately but then the two will eventually drift and your documentaiton will be a lie. Valdo solves all these problems: write validation once using a real programming language, use it everywhere.

Features:

* Mechanism to translate error messages.
* Out-of-the-box translations for some languages.
* Supports the latest JSON Schema specification (2020-12).
* Pure Go.
* No code generation, no reflection, no unsafe code.
* User-freindly error messages.
* Concurrency-safe, no global state.
* Strict by default, without imlicit type casting.
* Type-safe, thanks to generics.

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
user, err := valdo.Unmarshal[User](validator, input)

// generate JSON Schema
schema := valdo.Schema(validator)
```

See [documentation](https://pkg.go.dev/github.com/orsinium-labs/valdo/valdo) for more.
