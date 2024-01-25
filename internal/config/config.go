package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
