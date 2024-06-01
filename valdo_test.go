package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo"
)

func TestNewFieldChecks(t *testing.T) {
	_ = valdo.NewFieldChecks[int](valdo.Positive[int]())
	_ = valdo.NewFieldChecks(valdo.GTE(0))
}
