package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	Data   url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: errors(map[string][]string{}),
	}
}

func (f *Form) HasField(field string, r *http.Request) bool {
	fieldValue := r.Form.Get(field)
	return fieldValue != ""
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	fieldValue := r.Form.Get(field)
	if len(fieldValue) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Data.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
