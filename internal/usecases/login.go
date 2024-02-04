package usecases

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) Login(res http.ResponseWriter, req *http.Request) {
	_ = r.Config.Session.RenewToken(req.Context())

	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := req.Form.Get("email")
	password := req.Form.Get("password")

	form := forms.New(req.Form)
	form.Required("email", "password")
	form.IsEmail("email")
	if !form.Valid() {
		render.RenderTemplate(res, req, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	id, _, err := r.Database.GetUserByEmailAndPassword(email, password)
	if err != nil {
		log.Println(err)
		r.Config.Session.Put(req.Context(), "error", "Invalid login credentials")
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}

	r.Config.Session.Put(req.Context(), "user_id", id)
	r.Config.Session.Put(req.Context(), "flash", "Logged in successfully!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
