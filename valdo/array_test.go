package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestArray_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.A(valdo.Int(valdo.Min(0)), valdo.MinItems(1))
	noErr(valdo.Validate(val, []byte(`[1, 3, 4]`)))
	isErr[valdo.ErrIndex](valdo.Validate(val, []byte(`[1, -3, 4]`)))
	isErr[valdo.ErrIndex](valdo.Validate(val, []byte(`["aragorn"]`)))
	isErr[valdo.ErrMinItems](valdo.Validate(val, []byte(`[]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`"aragorn"`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`123`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
}

func TestArray_Schema(t *testing.T) {
	t.Parallel()
	{
		val := valdo.A(valdo.Int(valdo.Min(0)), valdo.MinItems(1))
		res := string(valdo.Schema(val))
		exp := `{"type":"array","items":{"type":"integer","minimum":0},"minItems":1}`
		isEq(res, exp)
	}
}
