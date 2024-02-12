package usecases

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the MakeReservation page
func (repo *Repository) MakeReservation(res http.ResponseWriter, req *http.Request) {
	log.Printf("[MakeReservation] passing through make reservation page endpoint")

	reservation, ok := repo.Config.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Printf("[MakeReservation] can't get reservation from session")
		repo.Config.Session.Put(req.Context(), "error", "Can't get reservation from session")
		http.Redirect(res, req, "/search-availability", http.StatusSeeOther)
		return
	}

	room, err := repo.Database.GetRoomByID(reservation.RoomID)
	if err != nil {
		log.Println(err)
		repo.Config.Session.Put(req.Context(), "error", "Can't get room from the reservation")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	reservation.Room.Name = room.Name
	reservation.Room.Description = room.Description

	repo.Config.Session.Put(req.Context(), "reservation", reservation)

	data := make(map[string]interface{})
	data["reservation"] = reservation

	dateMap := make(map[string]string)
	dateMap["start_date"] = reservation.StartDate.Format("2006-01-02")
	dateMap["end_date"] = reservation.EndDate.Format("2006-01-02")

	log.Printf("[MakeReservation] rendering make reservation page with dateMap: %v", dateMap)

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
