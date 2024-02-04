package usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/go-chi/chi"
)

// Responsible for rendering the Choose room page
func (repo *Repository) ChooseRoom(res http.ResponseWriter, req *http.Request) {
	pathID := chi.URLParam(req, "id")
	log.Printf("Receiving room reservation request for Room ID => %v", pathID)
	roomID, err := strconv.Atoi(pathID)
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	reservation, ok := repo.Config.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(res, err)
		return
	}

	reservation.RoomID = roomID
	repo.Config.Session.Put(req.Context(), "reservation", res)

	http.Redirect(res, req, "/make-reservation", http.StatusSeeOther)
}
