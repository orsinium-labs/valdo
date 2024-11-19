package valdo

type ArrayType struct {
	elem Validator
}

func Array(elem Validator) ArrayType {
	return ArrayType{
		elem: elem,
	}
}

func (a ArrayType) Validate(data any) Error {
	switch d := data.(type) {
	case []any:
		return a.validateArray(d)
	default:
		return ErrType{Got: getTypeName(data), Expected: "array"}
	}
}

func (a ArrayType) validateArray(data []any) Error {
	for i, val := range data {
		err := a.elem.Validate(val)
		if err != nil {
			return ErrIndex{Index: i, Err: err}
		}
	}
	return nil
}

func getTypeName(v any) string {
	if v == nil {
		return "null"
	}
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return "integer"
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return "unsigned integer"
	case bool:
		return "boolean"
	case string:
		return "string"
	case float32, float64:
		return "number"
	default:
		return ""
	}
}
