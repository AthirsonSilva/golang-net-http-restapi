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
	log.Printf("[ChooseRoom] passing through room choosing endpoint => %v", req.URL.Path)

	pathVar := helpers.GetPathVariableFromRequest(req)

	log.Printf("[ChooseRoom] parsing path variables from request %v", pathVar)

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
	repo.Config.Session.Put(req.Context(), "reservation", res)

	http.Redirect(res, req, "/make-reservation", http.StatusSeeOther)
}
