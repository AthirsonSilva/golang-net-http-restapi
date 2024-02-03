package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (repo *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	// Responsible for rendering the Contact page
	render.RenderTemplate(res, req, "contact.page.tmpl", &models.TemplateData{})
}
