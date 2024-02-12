package usecases

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) AdminAllReservations(res http.ResponseWriter, req *http.Request) {
	reservations, err := r.Database.GetAllReservations()
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Error getting all reservations", "admin-all-reservations")
		return
	}

	data := make(map[string]interface{})
	data["reservations"] = reservations

	render.RenderTemplate(res, req, "admin-all-reservations.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
