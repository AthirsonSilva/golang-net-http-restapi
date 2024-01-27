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
func (form *Form) HasField(field string) bool {
	hasField := form.Data.Get(field)
	return hasField != ""
}

// Checks if the form is valid
func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

// Verifies all the required fields
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Checks if the given field from the request has the passed minimum lenght
func (form *Form) MinLength(field string, length int, request *http.Request) bool {
	fieldValue := request.Form.Get(field)
	if len(fieldValue) < length {
		form.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// Validates if the given email is a valid one
func (form *Form) IsEmail(field string) {
	email := form.Data.Get(field)
	if !govalidator.IsEmail(email) {
		form.Errors.Add(field, "Invalid email address")
	}
}
