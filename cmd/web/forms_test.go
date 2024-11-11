package main

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_Form_Has(t *testing.T) {
	form := NewForm(nil)

	if form.Has("a") {
		t.Error("Form is expected to not have any fields")
	}

	vals := url.Values{}
	vals.Add("a", "a")
	vals.Add("b", "b")

	form = NewForm(vals)

	if !form.Has("a") {
		t.Error("Form is expected to have field \"a\", but method <Has> says there is no such field")
	}
	if !form.Has("b") {
		t.Error("Form is expected to have field \"b\", but method <Has> says there is no such field")
	}
}

func Test_Form_Required(t *testing.T) {
	req := httptest.NewRequest("POST", "/testing", nil)
	form := NewForm(req.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("Form doesn't have required field, but somehow considered valid")
	}

	postedData := url.Values{}
	postedData.Set("a", "a")
	postedData.Set("b", "b")
	postedData.Set("c", "c")

	form = NewForm(postedData)

	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("Form have all the required fields, but somehow considered invalid")
	}
}

func Test_Form_Check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "pwd", "password is required")

	if form.Valid() {
		t.Error("Form is valid but has errors by test logic")
	}
}

func Test_Form_Error_Get(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "pwd", "password is required")

	errMsg := form.Errors.Get("pwd")

	if len(errMsg) == 0 {
		t.Error("Should have error, but none is present")
	}

	errMsg = form.Errors.Get("blabla")
	if len(errMsg) != 0 {
		t.Error("Has error for unchecked field")
	}
}
