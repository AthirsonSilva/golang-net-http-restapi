package usecases

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) Login(res http.ResponseWriter, req *http.Request) {
	_ = r.Config.Session.RenewToken(req.Context())

	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		log.Println(err)
		r.Config.Session.Put(req.Context(), "error", "Cannot parse form")
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
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

	id, hashedPassword, err := r.Database.GetUserByEmailAndPassword(email, password)
	if err != nil || id == 0 {
		log.Println(err)
		RedirectWithError(r, req, res, "Invalid email or password", "/login")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		RedirectWithError(r, req, res, "Invalid email or password", "/login")
		return
	} else if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Invalid email or password", "/login")
		return
	}

	r.Config.Session.Put(req.Context(), "user_id", id)
	r.Config.Session.Put(req.Context(), "flash", "Logged in successfully!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
