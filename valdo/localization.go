package valdo

import "github.com/orsinium-labs/jsony"

// Locales maps language code to [Locale].
//
// It can [Wrap] a [Validator] to translate error messages to the selected language.
type Locales map[string]Locale

type Locale map[Error]string

func (ls Locales) Wrap(lang string, v Validator) Validator {
	locale, hasLang := ls[lang]
	if !hasLang {
		return v
	}
	return locale.Wrap(v)
}

func (loc Locale) Wrap(v Validator) Validator {
	return locVal{v: v, loc: loc}
}

type locVal struct {
	v   Validator
	loc Locale
}

// Valdiate implements [Validator].
func (lv locVal) Validate(data any) Error {
	err := lv.v.Validate(data)
	if err != nil {
		return lv.translate(err)
	}
	return nil
}

// translate the given error message.
func (lv locVal) translate(err Error) Error {
	format, found := lv.loc[err.GetDefault()]
	if !found {
		return err
	}
	return err.SetFormat(format)
}

// Schema implements [Validator].
func (lv locVal) Schema() jsony.Object {
	return lv.v.Schema()
}