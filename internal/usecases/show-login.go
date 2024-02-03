package usecases

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) LoginPage(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "login.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}
