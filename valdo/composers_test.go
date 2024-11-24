package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestAllOf(t *testing.T) {
	t.Parallel()
	val := valdo.AllOf(
		valdo.Int(valdo.Min(2)),
		valdo.Int(valdo.Max(4)),
	)
	noErr(valdo.Validate(val, []byte(`2`)))
	noErr(valdo.Validate(val, []byte(`3`)))
	noErr(valdo.Validate(val, []byte(`4`)))
	isErr[valdo.ErrMin](valdo.Validate(val, []byte(`1`)))
	isErr[valdo.ErrMax](valdo.Validate(val, []byte(`5`)))
}

func TestAnyOf(t *testing.T) {
	t.Parallel()
	val := valdo.AnyOf(
		valdo.Int(valdo.Min(5)),
		valdo.Int(valdo.Max(2)),
	)
	noErr(valdo.Validate(val, []byte(`1`)))
	noErr(valdo.Validate(val, []byte(`2`)))
	noErr(valdo.Validate(val, []byte(`5`)))
	noErr(valdo.Validate(val, []byte(`5`)))
	isErr[valdo.ErrAnyOf](valdo.Validate(val, []byte(`3`)))
	isErr[valdo.ErrAnyOf](valdo.Validate(val, []byte(`4`)))
}

func TestNot(t *testing.T) {
	t.Parallel()
	val := valdo.Not(
		valdo.Int(valdo.Min(4)),
	)
	noErr(valdo.Validate(val, []byte(`1`)))
	noErr(valdo.Validate(val, []byte(`2`)))
	noErr(valdo.Validate(val, []byte(`3`)))
	isErr[valdo.ErrNot](valdo.Validate(val, []byte(`4`)))
	isErr[valdo.ErrNot](valdo.Validate(val, []byte(`5`)))
}
