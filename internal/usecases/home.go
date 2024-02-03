package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the Home page
func (repo *Repository) Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "home.page.tmpl", &models.TemplateData{})
}
