package valdo

import (
	"github.com/orsinium-labs/jsony"
)

type Validator interface {
	Validate(data any) Errors
	Schema() jsony.Object
}
