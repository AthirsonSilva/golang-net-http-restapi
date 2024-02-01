package usecases

import (
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/go-chi/chi"
)

// Responsible for rendering the Choose room page
func (repo *Repository) ChooseRoom(responseWriter http.ResponseWriter, request *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	res, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(responseWriter, err)
		return
	}

	res.RoomID = roomID
	repo.Config.Session.Put(request.Context(), "reservation", res)

	http.Redirect(responseWriter, request, "/make-reservation", http.StatusSeeOther)
}
