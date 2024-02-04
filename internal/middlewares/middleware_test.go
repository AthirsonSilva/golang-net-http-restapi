package middlewares

import (
	"net/http"
	"testing"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
)

func TestNoSurf(t *testing.T) {
	var handler config.HttpHandler
	h := NoSurf(&handler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("Type is not http.Handler, but is %T", v)
	}
}

func TestSessionLoads(t *testing.T) {
	var handler config.HttpHandler
	h := SessionLoad(&handler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("Type is not http.Handler, but is %T", v)
	}
}
