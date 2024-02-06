package usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for receiving the data from the Availability page
func (repo *Repository) PostAvailability(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	startDate := helpers.ConvertDateFromString(req.Form.Get("start"), res)
	endDate := helpers.ConvertDateFromString(req.Form.Get("end"), res)

	rooms, err := Repo.Database.SearchAvailabilityByDateForAllRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	for _, room := range rooms {
		log.Printf("Found room: %v", room.Name)
	}

	if len(rooms) == 0 {
		Repo.Config.Session.Put(req.Context(), "error", "No availability")
		http.Redirect(res, req, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	reservation := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    0,
	}

	repo.Config.Session.Put(req.Context(), "reservation", reservation)

	render.RenderTemplate(res, req, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Responsible for rendering the Availability JSON page
func (repo *Repository) PostAvailabilityJSON(res http.ResponseWriter, req *http.Request) {
	startDate := helpers.ConvertDateFromString(req.Form.Get("start"), res)
	endDate := helpers.ConvertDateFromString(req.Form.Get("end"), res)

	roomID, err := strconv.Atoi(req.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(res, err)
	}

	available, err := repo.Database.SearchAvailabilityByDateAndRoom(startDate, endDate, roomID)
	helpers.JsonResponse(res, http.StatusOK, available)
}
