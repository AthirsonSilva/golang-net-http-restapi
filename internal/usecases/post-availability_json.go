package usecases

import (
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
)

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
