package usecases

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for receiving the data from the Availability page
func (repo *Repository) PostAvailability(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	startDate := helpers.ConvertDateFromString(request.Form.Get("start"), responseWriter)
	endDate := helpers.ConvertDateFromString(request.Form.Get("end"), responseWriter)

	rooms, err := Repo.Database.SearchAvailabilityByDateForAllRooms(startDate, endDate)
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
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    0,
	}

	repo.Config.Session.Put(request.Context(), "reservation", res)

	render.RenderTemplate(responseWriter, request, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
