package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestBool_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Bool()
	noErr(valdo.Validate(val, []byte(`true`)))
	noErr(valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`1`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`0`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`""`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`{}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`[]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
	isErr[valdo.ErrNoInput](valdo.Validate(val, []byte(``)))
}

func TestInt_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Int()
	noErr(valdo.Validate(val, []byte(`1`)))
	noErr(valdo.Validate(val, []byte(`1.0`)))
	noErr(valdo.Validate(val, []byte(`14`)))
	noErr(valdo.Validate(val, []byte(`-14`)))
	noErr(valdo.Validate(val, []byte(`-14.0`)))
	noErr(valdo.Validate(val, []byte(`0`)))
	noErr(valdo.Validate(val, []byte(`0.0`)))
	noErr(valdo.Validate(val, []byte(`-0.0`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`3.4`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`0.1`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`""`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`{}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
	isErr[valdo.ErrNoInput](valdo.Validate(val, []byte(``)))
}

func TestFloat64_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Float64()
	noErr(valdo.Validate(val, []byte(`1`)))
	noErr(valdo.Validate(val, []byte(`1.0`)))
	noErr(valdo.Validate(val, []byte(`14`)))
	noErr(valdo.Validate(val, []byte(`-14`)))
	noErr(valdo.Validate(val, []byte(`-14.0`)))
	noErr(valdo.Validate(val, []byte(`0`)))
	noErr(valdo.Validate(val, []byte(`0.0`)))
	noErr(valdo.Validate(val, []byte(`-0.0`)))
	noErr(valdo.Validate(val, []byte(`3.4`)))
	noErr(valdo.Validate(val, []byte(`-3.4`)))
	noErr(valdo.Validate(val, []byte(`-0.1`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`""`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`{}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
	isErr[valdo.ErrNoInput](valdo.Validate(val, []byte(``)))
}

func TestString_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.String()
	noErr(valdo.Validate(val, []byte(`""`)))
	noErr(valdo.Validate(val, []byte(`"hi"`)))
	noErr(valdo.Validate(val, []byte(`"hello"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`1`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`1.4`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`[]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`{}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`["hi"]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
	isErr[valdo.ErrNoInput](valdo.Validate(val, []byte(``)))
}

func TestNull_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Null()
	noErr(valdo.Validate(val, []byte(`null`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`1`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`1.4`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`[]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`{}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`["hi"]`)))
	isErr[valdo.ErrNoInput](valdo.Validate(val, []byte(``)))
}

func TestAny_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Any()
	noErr(valdo.Validate(val, []byte(`null`)))
	noErr(valdo.Validate(val, []byte(`false`)))
	noErr(valdo.Validate(val, []byte(`1`)))
	noErr(valdo.Validate(val, []byte(`1.4`)))
	noErr(valdo.Validate(val, []byte(`[]`)))
	noErr(valdo.Validate(val, []byte(`{}`)))
	noErr(valdo.Validate(val, []byte(`"hi"`)))
	noErr(valdo.Validate(val, []byte(`["hi"]`)))
	isErr[valdo.ErrNoInput](valdo.Validate(val, []byte(``)))
}

func TestConst_String_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Const("hi")
	noErr(valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`""`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`"hello"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`13`)))

	val = valdo.Const("42")
	noErr(valdo.Validate(val, []byte(`"42"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`42`)))
}

func TestConst_Bool_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Const(true)
	noErr(valdo.Validate(val, []byte(`true`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"true"`)))

	val = valdo.Const(false)
	noErr(valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`true`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"false"`)))
}

func TestConst_Int_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.Const(42)
	noErr(valdo.Validate(val, []byte(`42`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`13`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"42"`)))
}

func TestPrimitive_Schema(t *testing.T) {
	t.Parallel()
	isEq(string(valdo.Schema(valdo.Bool())), `{"type":"boolean"}`)
	isEq(string(valdo.Schema(valdo.String())), `{"type":"string"}`)
	isEq(string(valdo.Schema(valdo.Int())), `{"type":"integer"}`)
	isEq(string(valdo.Schema(valdo.Float64())), `{"type":"number"}`)
	isEq(string(valdo.Schema(valdo.Null())), `{"type":"null"}`)
	isEq(string(valdo.Schema(valdo.Any())), `{}`)
}

func TestConst_Schema(t *testing.T) {
	isEq(string(valdo.Schema(valdo.Const("hi"))), `{"const":"hi"}`)
	isEq(string(valdo.Schema(valdo.Const(true))), `{"const":true}`)
	isEq(string(valdo.Schema(valdo.Const(13))), `{"const":13}`)
}
