package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form is a wrapper for a url.Values object
type Form struct {
	Data   url.Values
	Errors errors
}

// Creates a new Form instance
func New(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: errors(map[string][]string{}),
	}
}

// Checks if the Form type has a value for the given field
func (f *Form) HasField(field string, r *http.Request) bool {
	fieldValue := r.Form.Get(field)
	return fieldValue != ""
}

// Checks if the form is valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Verifies all the required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Checks if the given field from the request has the passed minimum lenght
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	fieldValue := r.Form.Get(field)
	if len(fieldValue) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// Validates if the given email is a valid one
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Data.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
