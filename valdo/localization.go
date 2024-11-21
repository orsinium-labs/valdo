package valdo

import "github.com/orsinium-labs/jsony"

type Locales map[string]Locale

type Locale map[Error]string

func (ls Locales) Wrap(v Validator) Localizator {
	return Localizator{v: v, ls: ls}
}

type Localizator struct {
	v  Validator
	ls Locales
}

func (lz Localizator) HasLang(lang string) bool {
	_, hasLang := lz.ls[lang]
	return hasLang
}

func (lz Localizator) Translate(lang string) Validator {
	locale, hasLang := lz.ls[lang]
	if !hasLang {
		return lz.v
	}
	return locVal{v: lz.v, loc: locale}
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
