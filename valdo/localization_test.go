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

func TestTranslate_Recursive(t *testing.T) {
	t.Parallel()
	locales := valdo.Locales{
		"ru-RU": valdo.Locale{
			valdo.ErrProperty{}: "в поле {name}: {error}",
			valdo.ErrType{}:     "значение должно иметь тип {expected}",
		},
	}
	origV := valdo.Object(
		valdo.P("items", valdo.A(valdo.Int())),
	)
	val := locales.Wrap("ru-RU", origV)
	noErr(valdo.Validate(val, []byte(`{"items": [1, 2, 3]}`)))
	exp := "в поле items: at 1: значение должно иметь тип integer"
	isEq(valdo.Validate(val, []byte(`{"items": [1, "hi", 3]}`)).Error(), exp)
}
