package valdo

import (
	"encoding/json"

	"github.com/orsinium-labs/jsony"
)

type Validator interface {
	Validate(data any) Error
	Schema() jsony.Object
}

func Schema(v Validator) []byte {
	return jsony.EncodeBytes(v.Schema())
}

// Read the input JSON, validate it, and put the data into the given target.
func Unmarshal[T any](v Validator, input []byte, target *T) error {
	err := Validate(v, input)
	if err != nil {
		return err
	}
	return json.Unmarshal(input, target)
}

func Validate(v Validator, input []byte) Error {
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
