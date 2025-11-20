package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestTuple_Validate(t *testing.T) {
	t.Parallel()
	val := valdo.T(valdo.S(), valdo.I())
	noErr(valdo.Validate(val, []byte(`["aragorn", 82]`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`{"name": "aragorn"}`)))
	isErr[valdo.ErrType](valdo.Validate(val, []byte(`null`)))
	isErr[valdo.ErrMinItems](valdo.Validate(val, []byte(`["aragorn"]`)))
	isErr[valdo.ErrMaxItems](valdo.Validate(val, []byte(`["aragorn", 82, 42]`)))
	isErr[valdo.ErrIndex](valdo.Validate(val, []byte(`["aragorn", "82"]`)))
	isErr[valdo.ErrIndex](valdo.Validate(val, []byte(`[14, 82]`)))
	isErr[valdo.ErrIndex](valdo.Validate(val, []byte(`[14, "aragorn"]`)))
}

func TestTuple_Schema(t *testing.T) {
	t.Parallel()
	val := valdo.T(valdo.S(), valdo.I())
	res := string(valdo.Schema(val))
	exp := `{"type":"array","items":false,"prefixItems":[{"type":"string"},{"type":"integer"}]}`
	isEq(res, exp)
}
