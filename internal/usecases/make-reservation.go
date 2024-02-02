package usecases

import (
	"errors"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the MakeReservation page
func (repo *Repository) MakeReservation(responseWriter http.ResponseWriter, request *http.Request) {
	res, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(responseWriter, errors.New("cannot get reservation from session"))
		return
	}

	room, err := repo.Database.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	res.Room.RoomName = room.RoomName

	repo.Config.Session.Put(request.Context(), "reservation", res)

	data := make(map[string]interface{})
	data["reservation"] = res

	dateMap := make(map[string]string)
	dateMap["start_date"] = res.StartDate.Format("2006-01-02")
	dateMap["end_date"] = res.EndDate.Format("2006-01-02")

	render.RenderTemplate(responseWriter, request, "make-reservation.page.tmpl", &models.TemplateData{
		Form:    forms.New(nil),
		Data:    data,
		DateMap: dateMap,
	})
}
