package usecases

import (
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
)

// Responsible for rendering the Availability JSON page
func (repo *Repository) PostAvailabilityJSON(responseWriter http.ResponseWriter, request *http.Request) {
	startDate := helpers.ConvertDateFromString(request.Form.Get("start"), responseWriter)
	endDate := helpers.ConvertDateFromString(request.Form.Get("end"), responseWriter)

	roomID, err := strconv.Atoi(request.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(responseWriter, err)
	}

	available, err := repo.Database.SearchAvailabilityByDateAndRoom(startDate, endDate, roomID)
	helpers.JsonResponse(responseWriter, http.StatusOK, available)
}
