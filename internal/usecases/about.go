package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the About page
func (repo *Repository) About(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "about.page.tmpl", &models.TemplateData{})
}
