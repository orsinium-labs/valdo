package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestStringConst_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.StringConst("hi")
	noErr(valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`""`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`"hello"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`13`)))

	val = valdo.StringConst("42")
	noErr(valdo.Validate(val, []byte(`"42"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`42`)))
}

func TestBoolConst_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.BoolConst(true)
	noErr(valdo.Validate(val, []byte(`true`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"true"`)))

	val = valdo.BoolConst(false)
	noErr(valdo.Validate(val, []byte(`false`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`true`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"false"`)))
}

func TestConst_Int_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.IntConst(42)
	noErr(valdo.Validate(val, []byte(`42`)))
	isErr[valdo.ErrConst](valdo.Validate(val, []byte(`13`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"42"`)))
}

func TestConst_Schema(t *testing.T) {
	isEq(string(valdo.Schema(valdo.StringConst("hi"))), `{"const":"hi"}`)
	isEq(string(valdo.Schema(valdo.BoolConst(true))), `{"const":true}`)
	isEq(string(valdo.Schema(valdo.IntConst(13))), `{"const":13}`)
}
