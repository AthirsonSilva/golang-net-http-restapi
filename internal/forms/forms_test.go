package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	request := httptest.NewRequest("POST", "/whatever", nil)
	form := New(request.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	request := httptest.NewRequest("POST", "/whatever", nil)
	form := New(request.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	request = httptest.NewRequest("POST", "/whatever", nil)
	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.HasField("whatever")
	if has {
		t.Error("This field should not exist")
	}

	postedData = url.Values{}
	postedData.Add("field", "field")

	form = New(postedData)
	has = form.HasField("field")
	if !has {
		t.Error("This field should exist")
	}
}

func TestForm_MinLength(t *testing.T) {
	request := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("short", 10, request)
	if form.Valid() {
		t.Error("Form shows min length for non-existent field")
	}
	isError := form.Errors.Get("short")
	if isError == "" {
		t.Error("Should have an error, but did not get one")
	}
	postedData = url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLength("some_field", 100, request)
	if form.Valid() {
		t.Error("Shows min length of 100 met when data is shorter")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("Should have an invalid email error")
	}

	postedData = url.Values{}
	postedData.Add("email", "james.monroe@examplepetstore.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Should not have an invalid email error")

	}

	postedData = url.Values{}
	postedData.Add("email", "x")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Should have a invalid email error")
	}
}
