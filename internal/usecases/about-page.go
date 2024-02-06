package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the About page
func (repo *Repository) About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "about.page.tmpl", &models.TemplateData{})
}
