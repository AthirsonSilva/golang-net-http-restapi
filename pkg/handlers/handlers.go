package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/render"
)

var Repo *Repository

type Repository struct {
	Config *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		Config: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.Config.Session.Put(r.Context(), "remote_ip", remoteIP)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.Config.Session.GetString(r.Context(), "remote_ip")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIP
	stringMap["test"] = "Data from the server"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
