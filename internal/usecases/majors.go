package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for rendering the Reservation Summary page
func (repo *Repository) Majors(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "majors.page.tmpl", &models.TemplateData{})
}
