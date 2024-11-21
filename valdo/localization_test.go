package valdo_test

import (
	"testing"

	"github.com/orsinium-labs/valdo/valdo"
)

func TestTranslate(t *testing.T) {
	t.Parallel()
	locales := valdo.Locales{
		"ru-RU": valdo.Locale{
			valdo.ErrMin{}: "значение должно быть не меньше {value}",
		},
	}
	origV := valdo.Int(valdo.Min(2), valdo.Max(8))

	// supported language
	{
		val := locales.Wrap("ru-RU", origV)
		noErr(valdo.Validate(val, []byte(`4`)))
		isEq(valdo.Validate(val, []byte(`1`)).Error(), "значение должно быть не меньше 2")
		isEq(string(valdo.Schema(val)), string(valdo.Schema(origV)))
	}

	// unknown language
	{
		val := locales.Wrap("nl-BE", origV)
		noErr(valdo.Validate(val, []byte(`4`)))
		isEq(valdo.Validate(val, []byte(`1`)).Error(), "must be greater than or equal to 2")
		isEq(string(valdo.Schema(val)), string(valdo.Schema(origV)))
	}

	// unknown message
	{
		val := locales.Wrap("ru-RU", origV)
		noErr(valdo.Validate(val, []byte(`4`)))
		isEq(valdo.Validate(val, []byte(`10`)).Error(), "must be less than or equal to 8")
	}

}
