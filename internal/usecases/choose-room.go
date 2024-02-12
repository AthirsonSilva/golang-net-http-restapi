package usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

// Responsible for rendering the Choose room page
func (repo *Repository) ChooseRoom(res http.ResponseWriter, req *http.Request) {
	pathVar := helpers.PathVar(req)
	roomID, err := strconv.Atoi(pathVar)
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
	reservation.UserID = repo.Config.Session.GetInt(req.Context(), "user_id")
	repo.Config.Session.Put(req.Context(), "reservation", reservation)

	log.Println("reservation: ", reservation)

	http.Redirect(res, req, "/make-reservation", http.StatusSeeOther)
}
