package valdo

import (
	"encoding/json"
	"errors"

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
	T = Tuple
)

type Validator interface {
	Validate(data any) Error
	Schema() jsony.Object
}

// Schema generates JSON Schema for the validator.
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
func Validate(v Validator, input []byte) error {
	if len(input) == 0 {
		return ErrNoInput{}
	}
	var data any
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}
	if countJSONMembers(input) != countDecodedMembers(data) {
		return errors.New("duplicate JSON object names")
	}
	vErr := v.Validate(data)
	if vErr != nil {
		return vErr
	}
	return nil
}

// input must be valid JSON.
func countJSONMembers(input []byte) int {
	count := 0
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case ':':
			count++
		case '"':
			for i++; input[i] != '"'; i++ {
				if input[i] == '\\' {
					i++
				}
			}
		}
	}
	return count
}

func countDecodedMembers(data any) int {
	count := 0
	switch data := data.(type) {
	case map[string]any:
		count = len(data)
		for _, value := range data {
			count += countDecodedMembers(value)
		}
	case []any:
		for _, value := range data {
			count += countDecodedMembers(value)
		}
	}
	return count
}
