package usecases

import (
	"errors"
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (repo *Repository) ReservationSummary(res http.ResponseWriter, req *http.Request) {
	log.Println("[ReservationSummary] reservation-summary")

	reservation, ok := repo.Config.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(res, errors.New("cannot get reservation"))
		return
	}

	log.Printf("[ReservationSummary] reservation: %v", reservation.Room.Name)

	data := make(map[string]interface{})
	data["reservation"] = reservation

	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")
	dateMap := make(map[string]string)
	dateMap["start_date"] = startDate
	dateMap["end_date"] = endDate

	log.Printf("[ReservationSummary] dateMap: %v", dateMap)

	render.RenderTemplate(res, req, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:    data,
		DateMap: dateMap,
	})
}
