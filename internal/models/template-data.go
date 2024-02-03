package models

import "github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"

// TemplateData holds data sent from handlers to templates and vice-versa
type TemplateData struct {
	DateMap         map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated bool
}
