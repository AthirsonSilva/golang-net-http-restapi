package usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (r *Repository) AdminDeleteReservationByID(res http.ResponseWriter, req *http.Request) {
	reservationId, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Invalid reservation id", "/admin/reservations/all")
		return
	}

	err = r.Database.DeleteReservationByID(reservationId)
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Could not delete reservation", "/admin/reservations/all")
		return
	}

	http.Redirect(res, req, "/admin/reservations/all", http.StatusSeeOther)
}
