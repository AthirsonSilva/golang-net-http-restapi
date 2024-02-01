package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for receiving the data from the Availability page
func (repo *Repository) PostAvailability(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	layout := "2006-01-02"
	parsed_start_date, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	parsed_end_date, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	rooms, err := Repo.Database.SearchAvailabilityByDateForAllRooms(parsed_start_date, parsed_end_date)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	for _, room := range rooms {
		log.Printf("Found room: %v", room.RoomName)
	}

	if len(rooms) == 0 {
		Repo.Config.Session.Put(request.Context(), "error", "No availability")
		http.Redirect(responseWriter, request, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: parsed_start_date,
		EndDate:   parsed_end_date,
		RoomID:    0,
	}

	repo.Config.Session.Put(request.Context(), "reservation", res)

	render.RenderTemplate(responseWriter, request, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
