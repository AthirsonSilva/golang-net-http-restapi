package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for rendering the Reservation Summary page
func (repo *Repository) Generals(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "generals.page.tmpl", &models.TemplateData{})
}
