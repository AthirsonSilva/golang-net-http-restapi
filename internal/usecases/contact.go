package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (repo *Repository) Contact(responseWriter http.ResponseWriter, request *http.Request) {
	// Responsible for rendering the Contact page
	render.RenderTemplate(responseWriter, request, "contact.page.tmpl", &models.TemplateData{})
}
