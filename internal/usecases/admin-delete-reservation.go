package usecases

import (
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/go-chi/chi/v5"
)

func (r *Repository) AdminDeleteReservationByID(res http.ResponseWriter, req *http.Request) {
	reservationId, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		helpers.ServerError(res, err)
	}

	err = r.Database.DeleteReservationByID(reservationId)
	if err != nil {
		helpers.ServerError(res, err)
	}

	http.Redirect(res, req, "/admin/reservations/all", http.StatusSeeOther)
}
