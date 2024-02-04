package config

import (
	"log"
	"net/http"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application system-wide config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
}

// Mock HttpHandler for testing purposes
type HttpHandler struct{}

// Mock ServerHTTP for testing purposes
func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// Creates instances for both the application's system-wide config and Session manager
var (
	App     AppConfig
	Session *scs.SessionManager
)
