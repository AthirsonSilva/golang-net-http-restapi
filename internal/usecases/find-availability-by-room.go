package usecases

import (
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"github.com/go-chi/chi"
)

func (r *Repository) FindAvailabilityByRoom(res http.ResponseWriter, req *http.Request) {
	roomId, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	room, err := r.Database.GetRoomByID(roomId)
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	data := make(map[string]interface{})
	data["room"] = room

	render.RenderTemplate(res, req, "find-availability-by-room.page.tmpl", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}
