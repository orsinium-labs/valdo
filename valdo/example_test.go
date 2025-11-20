package valdo_test

import (
	"fmt"

	"github.com/orsinium-labs/valdo/valdo"
)

func ExampleDefaultLocales() {
	language := "nl"
	original := valdo.Int()
	translated := valdo.DefaultLocales.Wrap(language, original)
	input := []byte(`"hi"`)
	err := valdo.Validate(translated, input)
	fmt.Println(err)
	// Output: ongeldig type: kreeg string, verwachtte integer
}

func ExampleLocales() {
	language := "nl"
	original := valdo.Int()
	locales := valdo.Locales{
		"nl": valdo.Locale{
			valdo.ErrType{}: "ongeldig type: kreeg {got}, verwachtte {expected}",
		},
	}
	translated := locales.Wrap(language, original)
	input := []byte(`"hi"`)
	err := valdo.Validate(translated, input)
	fmt.Println(err)
	// Output: ongeldig type: kreeg string, verwachtte integer
}

func ExampleLocale() {
	original := valdo.Int()
	locale := valdo.Locale{
		valdo.ErrType{}: "ongeldig type: kreeg {got}, verwachtte {expected}",
	}
	translated := locale.Wrap(original)
	input := []byte(`"hi"`)
	err := valdo.Validate(translated, input)
	fmt.Println(err)
	// Output: ongeldig type: kreeg string, verwachtte integer
}

func ExampleString() {
	validator := valdo.String()
	input := []byte(`"hi"`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleInt() {
	validator := valdo.Int()
	input := []byte(`13`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleBool() {
	validator := valdo.Bool()
	input := []byte(`true`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleFloat64() {
	validator := valdo.Float64()
	input := []byte(`3.14`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleArray() {
	validator := valdo.Array(valdo.Int())
	input := []byte(`[3, 4, 5]`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleObject() {
	validator := valdo.Object(
		valdo.P("age", valdo.Int()),
	)
	input := []byte(`{"age": 42}`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleEnum() {
	validator := valdo.Enum("red", "green", "blue")
	input := []byte(`"green"`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleValidate() {
	validator := valdo.Object(
		valdo.P("age", valdo.Int()),
	)
	input := []byte(`{"age": 42}`)
	err := valdo.Validate(validator, input)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleSchema() {
	validator := valdo.Int()
	schema := valdo.Schema(validator)
	fmt.Println(string(schema))
	// Output: {"type":"integer"}
}
