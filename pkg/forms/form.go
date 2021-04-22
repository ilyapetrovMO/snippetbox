package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

func (f *Form) MinLength(field string, len int) {
	v := f.Get(field)
	if v == "" {
		return
	}

	if utf8.RuneCountInString(v) < len {
		f.Errors.Add(field, fmt.Sprintf("This field cannot be shorter than %d", len))
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	v := f.Get(field)
	if v == "" {
		return
	}

	if !pattern.MatchString(v) {
		f.Errors.Add(field, "This field is invalid")
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
