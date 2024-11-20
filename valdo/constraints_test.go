package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestMultipleOf(t *testing.T) {
	t.Parallel()
	val := valdo.Int(valdo.MultipleOf(3))
	noErr(valdo.Validate(val, []byte(`0`)))
	noErr(valdo.Validate(val, []byte(`3`)))
	noErr(valdo.Validate(val, []byte(`6`)))
	noErr(valdo.Validate(val, []byte(`-6`)))
	isErr[valdo.ErrMultipleOf](valdo.Validate(val, []byte(`1`)))
	isErr[valdo.ErrMultipleOf](valdo.Validate(val, []byte(`-1`)))
	isErr[valdo.ErrMultipleOf](valdo.Validate(val, []byte(`2`)))
	isErr[valdo.ErrMultipleOf](valdo.Validate(val, []byte(`-2`)))
	isErr[valdo.ErrMultipleOf](valdo.Validate(val, []byte(`10`)))
	isErr[valdo.ErrMultipleOf](valdo.Validate(val, []byte(`-10`)))

	isEq(string(valdo.Schema(val)), `{"type":"integer","multipleOf":3}`)
}

func TestMin(t *testing.T) {
	t.Parallel()
	val := valdo.Int(valdo.Min(3))
	noErr(valdo.Validate(val, []byte(`3`)))
	noErr(valdo.Validate(val, []byte(`4`)))
	noErr(valdo.Validate(val, []byte(`6`)))
	isErr[valdo.ErrMin](valdo.Validate(val, []byte(`2`)))
	isErr[valdo.ErrMin](valdo.Validate(val, []byte(`0`)))
	isErr[valdo.ErrMin](valdo.Validate(val, []byte(`-6`)))

	isEq(string(valdo.Schema(val)), `{"type":"integer","minimum":3}`)
}

func TestExclMin(t *testing.T) {
	t.Parallel()
	val := valdo.Int(valdo.ExclMin(3))
	noErr(valdo.Validate(val, []byte(`4`)))
	noErr(valdo.Validate(val, []byte(`6`)))
	isErr[valdo.ErrExclMin](valdo.Validate(val, []byte(`3`)))
	isErr[valdo.ErrExclMin](valdo.Validate(val, []byte(`2`)))
	isErr[valdo.ErrExclMin](valdo.Validate(val, []byte(`0`)))
	isErr[valdo.ErrExclMin](valdo.Validate(val, []byte(`-6`)))

	isEq(string(valdo.Schema(val)), `{"type":"integer","exclusiveMinimum":3}`)
}

func TestMax(t *testing.T) {
	t.Parallel()
	val := valdo.Int(valdo.Max(3))
	noErr(valdo.Validate(val, []byte(`3`)))
	noErr(valdo.Validate(val, []byte(`2`)))
	noErr(valdo.Validate(val, []byte(`-6`)))
	isErr[valdo.ErrMax](valdo.Validate(val, []byte(`4`)))
	isErr[valdo.ErrMax](valdo.Validate(val, []byte(`5`)))
	isErr[valdo.ErrMax](valdo.Validate(val, []byte(`16`)))

	isEq(string(valdo.Schema(val)), `{"type":"integer","maximum":3}`)
}

func TestExclMax(t *testing.T) {
	t.Parallel()
	val := valdo.Int(valdo.ExclMax(3))
	noErr(valdo.Validate(val, []byte(`2`)))
	noErr(valdo.Validate(val, []byte(`-6`)))
	isErr[valdo.ErrExclMax](valdo.Validate(val, []byte(`3`)))
	isErr[valdo.ErrExclMax](valdo.Validate(val, []byte(`4`)))
	isErr[valdo.ErrExclMax](valdo.Validate(val, []byte(`5`)))
	isErr[valdo.ErrExclMax](valdo.Validate(val, []byte(`16`)))

	isEq(string(valdo.Schema(val)), `{"type":"integer","exclusiveMaximum":3}`)
}

func TestMinLen(t *testing.T) {
	t.Parallel()
	val := valdo.String(valdo.MinLen(3))
	noErr(valdo.Validate(val, []byte(`"hel"`)))
	noErr(valdo.Validate(val, []byte(`"hell"`)))
	noErr(valdo.Validate(val, []byte(`"hello"`)))
	isErr[valdo.ErrMinLen](valdo.Validate(val, []byte(`"hi"`)))
	isErr[valdo.ErrMinLen](valdo.Validate(val, []byte(`"!"`)))
	isErr[valdo.ErrMinLen](valdo.Validate(val, []byte(`""`)))

	isEq(string(valdo.Schema(val)), `{"type":"string","minLength":3}`)
}

func TestMaxLen(t *testing.T) {
	t.Parallel()
	val := valdo.String(valdo.MaxLen(3))
	noErr(valdo.Validate(val, []byte(`"hel"`)))
	noErr(valdo.Validate(val, []byte(`"hi"`)))
	noErr(valdo.Validate(val, []byte(`"!"`)))
	noErr(valdo.Validate(val, []byte(`""`)))
	isErr[valdo.ErrMaxLen](valdo.Validate(val, []byte(`"hell"`)))
	isErr[valdo.ErrMaxLen](valdo.Validate(val, []byte(`"hello"`)))

	isEq(string(valdo.Schema(val)), `{"type":"string","maxLength":3}`)
}

func TestPattern(t *testing.T) {
	t.Parallel()
	val := valdo.String(valdo.Pattern("[AB]"))
	noErr(valdo.Validate(val, []byte(`"A"`)))
	noErr(valdo.Validate(val, []byte(`"B"`)))
	noErr(valdo.Validate(val, []byte(`"Bxx"`)))
	noErr(valdo.Validate(val, []byte(`"xxB"`)))
	noErr(valdo.Validate(val, []byte(`"xxBxx"`)))
	isErr[valdo.ErrPattern](valdo.Validate(val, []byte(`"xxxx"`)))
	isErr[valdo.ErrPattern](valdo.Validate(val, []byte(`"hello"`)))
	isErr[valdo.ErrPattern](valdo.Validate(val, []byte(`""`)))

	isEq(string(valdo.Schema(val)), `{"type":"string","pattern":"[AB]"}`)
}

func TestContains(t *testing.T) {
	t.Parallel()
	val := valdo.Array(valdo.Any(), valdo.Contains(valdo.Int()))
	noErr(valdo.Validate(val, []byte(`[1, 2, 3]`)))
	noErr(valdo.Validate(val, []byte(`["", 2]`)))
	noErr(valdo.Validate(val, []byte(`[2, ""]`)))
	noErr(valdo.Validate(val, []byte(`[2, "", 3]`)))
	isErr[valdo.ErrContains](valdo.Validate(val, []byte(`[]`)))
	isErr[valdo.ErrContains](valdo.Validate(val, []byte(`[""]`)))
	isErr[valdo.ErrContains](valdo.Validate(val, []byte(`["", 2.3]`)))
	isErr[valdo.ErrContains](valdo.Validate(val, []byte(`[true, 2.3]`)))
	isErr[valdo.ErrContains](valdo.Validate(val, []byte(`[true]`)))

	isEq(string(valdo.Schema(val)), `{"type":"array","contains":{"type":"integer"}}`)
}

func TestMinItems(t *testing.T) {
	t.Parallel()
	val := valdo.Array(valdo.Any(), valdo.MinItems(3))
	noErr(valdo.Validate(val, []byte(`[5, 6, 7]`)))
	noErr(valdo.Validate(val, []byte(`[5, 6, 7, 8]`)))
	isErr[valdo.ErrMinItems](valdo.Validate(val, []byte(`[]`)))
	isErr[valdo.ErrMinItems](valdo.Validate(val, []byte(`[5]`)))
	isErr[valdo.ErrMinItems](valdo.Validate(val, []byte(`[5, 6]`)))

	isEq(string(valdo.Schema(val)), `{"type":"array","minItems":3}`)
}

func TestMaxItems(t *testing.T) {
	t.Parallel()
	val := valdo.Array(valdo.Any(), valdo.MaxItems(3))
	noErr(valdo.Validate(val, []byte(`[]`)))
	noErr(valdo.Validate(val, []byte(`[6]`)))
	noErr(valdo.Validate(val, []byte(`[6, 7]`)))
	noErr(valdo.Validate(val, []byte(`[6, 7, 8]`)))
	isErr[valdo.ErrMaxItems](valdo.Validate(val, []byte(`[5, 6, 7, 8]`)))
	isErr[valdo.ErrMaxItems](valdo.Validate(val, []byte(`[5, 6, 7, 8, 9]`)))

	isEq(string(valdo.Schema(val)), `{"type":"array","maxItems":3}`)
}

func TestPropertyNames(t *testing.T) {
	t.Parallel()
	val := valdo.Map(nil, valdo.PropertyNames(valdo.MaxLen(2)))
	noErr(valdo.Validate(val, []byte(`{}`)))
	noErr(valdo.Validate(val, []byte(`{"hi": 1, "h": 2}`)))
	isErr[valdo.ErrPropertyNames](valdo.Validate(val, []byte(`{"hel": 1, "h": 2}`)))
	isErr[valdo.ErrPropertyNames](valdo.Validate(val, []byte(`{"h": 1, "hel": 2}`)))
	isErr[valdo.ErrPropertyNames](valdo.Validate(val, []byte(`{"hello": 1, "h": 2}`)))
	isErr[valdo.ErrPropertyNames](valdo.Validate(val, []byte(`{"hello": 1}`)))

	isEq(string(valdo.Schema(val)), `{"type":"object","propertyNames":{"maxLength":2}}`)
}

func TestMinProperties(t *testing.T) {
	t.Parallel()
	val := valdo.Map(nil, valdo.MinProperties(2))
	noErr(valdo.Validate(val, []byte(`{"a": 1, "b": 2}`)))
	noErr(valdo.Validate(val, []byte(`{"a": 1, "b": 2, "c": 3}`)))
	isErr[valdo.ErrMinProperties](valdo.Validate(val, []byte(`{"hel": 1}`)))
	isErr[valdo.ErrMinProperties](valdo.Validate(val, []byte(`{}`)))

	isEq(string(valdo.Schema(val)), `{"type":"object","minProperties":2}`)
}

func TestMaxProperties(t *testing.T) {
	t.Parallel()
	val := valdo.Map(nil, valdo.MaxProperties(2))
	noErr(valdo.Validate(val, []byte(`{"a": 1, "b": 2}`)))
	noErr(valdo.Validate(val, []byte(`{"a": 1}`)))
	noErr(valdo.Validate(val, []byte(`{}`)))
	isErr[valdo.ErrMaxProperties](valdo.Validate(val, []byte(`{"a": 1, "b": 2, "c": 3}`)))

	isEq(string(valdo.Schema(val)), `{"type":"object","maxProperties":2}`)
}
