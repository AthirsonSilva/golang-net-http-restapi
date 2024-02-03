package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) AdminDashboard(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "admin-dashboard.page.tmpl", &models.TemplateData{})
}
