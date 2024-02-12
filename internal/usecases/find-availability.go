package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for rendering the Availability page
func (r *Repository) Availability(res http.ResponseWriter, req *http.Request) {
	rooms, err := r.Database.GetAllRooms()
	if err != nil {
		helpers.ServerError(res, err)
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	render.RenderTemplate(res, req, "search-availability.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
