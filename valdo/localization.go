package valdo

import "github.com/orsinium-labs/jsony"

type Locales map[string]Locale

type Locale map[Error]string

func (ls Locales) Wrap(lang string, v Validator) Validator {
	locale, hasLang := ls[lang]
	if !hasLang {
		return v
	}
	return locVal{v: v, loc: locale}
}

type locVal struct {
	v   Validator
	loc Locale
}

func (lv locVal) Validate(data any) Error {
	err := lv.v.Validate(data)
	if err != nil {
		return lv.translate(err)
	}
	return nil
}

func (lv locVal) translate(err Error) Error {
	format, found := lv.loc[err.GetDefault()]
	if !found {
		return err
	}
	return err.SetFormat(format)
}

func (lv locVal) Schema() jsony.Object {
	return lv.v.Schema()
}
