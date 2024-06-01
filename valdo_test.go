package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo"
)

func TestNewFieldChecks(t *testing.T) {
	_ = valdo.F(valdo.Positive[int]())
	_ = valdo.F(valdo.GTE(0))
	_ = valdo.F(valdo.GTE(0), valdo.Positive[int]())
	_ = valdo.F(valdo.MaxLen[string, byte](4))
	_ = valdo.F(valdo.MaxLen[string, rune](4))
	_ = valdo.F(valdo.SMaxLen[[]rune](4))
	_ = valdo.F(valdo.SMaxLen[[]int](4))
}
