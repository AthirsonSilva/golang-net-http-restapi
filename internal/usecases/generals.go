package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for rendering the Reservation Summary page
func (repo *Repository) Generals(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "generals.page.tmpl", &models.TemplateData{})
}
