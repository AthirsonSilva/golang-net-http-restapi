package usecases

import (
	"errors"
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the MakeReservation page
func (repo *Repository) MakeReservation(res http.ResponseWriter, req *http.Request) {
	log.Printf("[MakeReservation] passing through make reservation page endpoint")

	reservation, ok := repo.Config.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(res, errors.New("cannot get reservation from session"))
		return
	}

	room, err := repo.Database.GetRoomByID(reservation.RoomID)
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	reservation.Room.Name = room.Name

	repo.Config.Session.Put(req.Context(), "reservation", res)

	data := make(map[string]interface{})
	data["reservation"] = reservation

	dateMap := make(map[string]string)
	dateMap["start_date"] = reservation.StartDate.Format("2006-01-02")
	dateMap["end_date"] = reservation.EndDate.Format("2006-01-02")

	render.RenderTemplate(
		res,
		req,
		"make-reservation.page.tmpl",
		&models.TemplateData{
			Form:    forms.New(nil),
			Data:    data,
			DateMap: dateMap,
		},
	)
}
