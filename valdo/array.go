package valdo

import "github.com/orsinium-labs/jsony"

// ArrayType is constructed by [Array].
type ArrayType struct {
	elem Validator
	cs   []Constraint[[]any]
}

func Array(elem Validator, cs ...Constraint[[]any]) ArrayType {
	return ArrayType{
		elem: elem,
	}.Constrain(cs...)
}

func (a ArrayType) Constrain(cs ...Constraint[[]any]) ArrayType {
	a.cs = append(a.cs, cs...)
	return a
}

// Validate implements [Validator].
func (a ArrayType) Validate(data any) Error {
	switch d := data.(type) {
	case []any:
		return a.validateArray(d)
	default:
		return ErrType{Got: getTypeName(data), Expected: "array"}
	}
}

func (a ArrayType) validateArray(data []any) Error {
	if data == nil {
		return ErrType{Got: "null", Expected: "array"}
	}
	res := Errors{}
	for i, val := range data {
		err := a.elem.Validate(val)
		if err != nil {
			res.Add(ErrIndex{Index: i, Err: err})
			break
		}
	}
	for _, c := range a.cs {
		res.Add(c.check(data))
	}
	return res.Flatten()
}

// Schema implements [Validator].
func (a ArrayType) Schema() jsony.Object {
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.SafeString("array")},
		jsony.Field{K: "items", V: a.elem.Schema()},
	}
	for _, c := range a.cs {
		res = append(res, c.field)
	}
	return res
}

func getTypeName(v any) string {
	if v == nil {
		return "null"
	}
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return "integer"
	case jsony.Int, jsony.Int8, jsony.Int16, jsony.Int32, jsony.Int64:
		return "integer"
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return "unsigned integer"
	case jsony.UInt, jsony.UInt8, jsony.UInt16, jsony.UInt32, jsony.UInt64:
		return "unsigned integer"
	case bool, jsony.Bool:
		return "boolean"
	case string, jsony.String:
		return "string"
	case float32, float64, jsony.Float32, jsony.Float64:
		return "number"
	case map[string]any, jsony.Object:
		return "object"
	case []any, jsony.MixedArray:
		return "array"
	default:
		return ""
	}
}
