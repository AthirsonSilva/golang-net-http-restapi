package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) AdminReservationsCalendar(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "admin-reservations-calendar.page.tmpl", &models.TemplateData{})
}
