package routes

import (
	"testing"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var appConfig config.AppConfig

	router := Routes(&appConfig)

	switch v := router.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("type is not *chi.Router, but is %T", v)
	}
}
