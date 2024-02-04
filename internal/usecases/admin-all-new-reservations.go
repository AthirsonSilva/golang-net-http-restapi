package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) AdminAllNewReservations(res http.ResponseWriter, req *http.Request) {
	reservations, err := r.Database.GetAllNewReservations()
	if err != nil {
		helpers.ServerError(res, err)
	}

	data := make(map[string]interface{})
	data["reservations"] = reservations

	render.RenderTemplate(res, req, "admin-all-reservations.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
