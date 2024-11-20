package valdo_test

import (
	"fmt"
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func noErr(e error) {
	if e != nil {
		panic(fmt.Sprintf("unexpected error: %T: %v", e, e))
	}
}

func isErr[T error](e error) {
	_, ok := e.(T)
	if !ok {
		panic(fmt.Sprintf("unexpected error type: %T: %v", e, e))
	}
}

func isEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("%v\n", a)
		fmt.Printf("%v\n", b)
		panic(fmt.Sprintf("%v != %v", a, b))
	}
}

func TestObject_Validate_Map(t *testing.T) {
	t.Parallel()
	val := valdo.O(
		valdo.P("name", valdo.S(valdo.MinLen(2))),
		valdo.P("admin", valdo.B()),
	)
	noErr(valdo.Validate(val, []byte(`{"name": "aragorn", "admin": true}`)))
	isErr[valdo.ErrRequired](valdo.Validate(val, []byte(`{"name": "aragorn"}`)))
	isErr[valdo.ErrProperty](valdo.Validate(val, []byte(`{"name": "", "admin": false}`)))
	isErr[valdo.ErrProperty](valdo.Validate(val, []byte(`{"name": 123, "admin": false}`)))
	isErr[valdo.ErrProperty](valdo.Validate(val, []byte(`{"name": "aragorn", "admin":""}`)))
	isErr[valdo.ErrUnexpected](valdo.Validate(val, []byte(`{"name":"aragorn","admin":true,"hi":1}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`["aragorn"]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"aragorn"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`123`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
	isErr[valdo.Errors](valdo.Validate(val, []byte(`{}`)))
}

func TestObject_Schema(t *testing.T) {
	t.Parallel()
	{
		val := valdo.O(
			valdo.P("name", valdo.S()),
		)
		res := string(valdo.Schema(val))
		exp := `{"type":"object","properties":{"name":{"type":"string"}},"required":["name"],"additionalProperties":false}`
		isEq(res, exp)
	}

	{
		val := valdo.O(
			valdo.P("name", valdo.S(valdo.MinLen(2))),
		)
		res := string(valdo.Schema(val))
		exp := `{"type":"object","properties":{"name":{"type":"string","minLength":2}},"required":["name"],"additionalProperties":false}`
		isEq(res, exp)
	}

	{
		val := valdo.O(
			valdo.P("name", valdo.S()).Optional(),
		)
		res := string(valdo.Schema(val))
		exp := `{"type":"object","properties":{"name":{"type":"string"}},"additionalProperties":false}`
		isEq(res, exp)
	}
}
