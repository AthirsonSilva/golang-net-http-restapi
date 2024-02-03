package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for rendering the Availability page
func (repo *Repository) Availability(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "search-availability.page.tmpl", &models.TemplateData{})
}
