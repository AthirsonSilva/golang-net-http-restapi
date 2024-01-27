package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var handler httpHandler
	h := NoSurf(&handler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("Type is not http.Handler, but is %T", v)
	}
}

func TestSessionLoads(t *testing.T) {
	var handler httpHandler
	h := SessionLoad(&handler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("Type is not http.Handler, but is %T", v)
	}
}
