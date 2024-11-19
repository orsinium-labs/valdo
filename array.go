package valdo

type Array struct {
	elem Validator
	msg  string
}

func A(elem Validator) Array {
	return Array{
		elem: elem,
		msg:  "invalid type: expected array, got %s",
	}
}

func (a Array) Validate(data any) Errors {
	switch d := data.(type) {
	case []any:
		return a.validateArray(d)
	case int, int8, int16, int32, int64:
		return newFieldError(a.msg, "integer")
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return newFieldError(a.msg, "unsigned integer")
	case bool:
		return newFieldError(a.msg, "boolean")
	case string:
		return newFieldError(a.msg, "string")
	case float32, float64:
		return newFieldError(a.msg, "number")
	default:
		return a.validateReflect(data)
	}
}

func (a Array) validateArray(data []any) Errors {
	for _, val := range data {
		err := a.elem.Validate(val)
		if err != nil {
			// TODO: include value index
			return err
		}
	}
	return nil
}

func (a Array) validateReflect(data any) Errors {
	panic("todo")
}
