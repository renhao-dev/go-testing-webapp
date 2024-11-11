package main

import "net/url"

type errors map[string][]string

func (e errors) Get(field string) string {
	fieldErrors := e[field]

	if len(fieldErrors) > 0 {
		return fieldErrors[0]
	}

	return ""
}

func (e errors) Add(field string, msg string) {
	e[field] = append(e[field], msg)
}

type Form struct {
	Values url.Values
	Errors errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		f.Check(f.Values.Get(field) != "", field, "Field can't be empty")
	}
}
func (f *Form) Has(field string) bool {
	return f.Values.Get(field) != ""
}
func (f *Form) Check(cond bool, field string, msg string) {
	if !cond {
		f.Errors.Add(field, msg)
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
