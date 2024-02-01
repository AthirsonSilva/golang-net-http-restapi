package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

// Responsible for rendering the Availability JSON page
func (repo *Repository) PostAvailabilityJSON(responseWriter http.ResponseWriter, request *http.Request) {
	response := models.JsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(response, "", "     ")
	if err != nil {
		helpers.ServerError(responseWriter, err)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(out)
}
