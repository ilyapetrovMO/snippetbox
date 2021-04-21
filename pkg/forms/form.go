package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, v := range fields {
		value := f.Get(v)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(v, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLength(field string, len int) {
	v := f.Get(field)
	if v == "" {
		return
	}

	if utf8.RuneCountInString(v) > len {
		f.Errors.Add(field, fmt.Sprintf("This field cannot be longer than %d", len))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	v := f.Get(field)
	if v == "" {
		return
	}

	for _, opt := range opts {
		if v == opt {
			return
		}
	}

	f.Errors.Add(field, "This field is invalid")
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
