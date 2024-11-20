package valdo

import (
	"encoding/json"

	"github.com/orsinium-labs/jsony"
)

// One-letter abbreviations for people living on the edge.
var (
	O = Object
	A = Array
	P = Property
	B = Bool
	S = String
	N = Null
	F = Float64
	I = Int
)

type Validator interface {
	Validate(data any) Error
	Schema() jsony.Object
}

func Schema(v Validator) []byte {
	return jsony.EncodeBytes(v.Schema())
}

// Read the input JSON, validate it, and unmarshal into the given type.
func Unmarshal[T any](v Validator, input []byte) (T, error) {
	var target T
	err := Validate(v, input)
	if err != nil {
		return target, err
	}
	err = json.Unmarshal(input, &target)
	return target, err
}

// Validate the given JSON.
func Validate(v Validator, input []byte) Error {
	if len(input) == 0 {
		return ErrNoInput{}
	}
	var data any
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}
	vErr := v.Validate(data)
	if vErr != nil {
		return vErr
	}
	return nil
}
