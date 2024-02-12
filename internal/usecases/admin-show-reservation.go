package usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"github.com/go-chi/chi/v5"
)

func (r *Repository) AdminShowSingleReservation(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "No data sent from client", "/admin/reservations-new")
		return
	}

	reservation, err := r.Database.GetReservationByID(id)
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Could not retrieve reservation", "/admin/reservation")
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(res, req, "admin-show-single-reservation.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
