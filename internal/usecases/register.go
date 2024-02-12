package usecases

import (
	"log"
	"net/http"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) Register(res http.ResponseWriter, req *http.Request) {
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
	firstName := req.Form.Get("first_name")
	lastName := req.Form.Get("last_name")
	phone := req.Form.Get("phone")
	password := req.Form.Get("password")

	form := forms.New(req.Form)
	form.Required("email", "password", "first_name", "last_name", "phone")
	form.MinLength(6, req, "password")
	form.IsEmail("email")
	if !form.Valid() {
		render.RenderTemplate(res, req, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	id, _, _ := r.Database.GetUserByEmailAndPassword(email, password)
	if id > 0 {
		RedirectWithError(r, req, res, "User already exists", "/login")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Cannot generate hashed password", "/login")
		return
	}

	password = string(hashedPassword)
	newUser := models.User{
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		Phone:       phone,
		Password:    password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		AccessLevel: 1,
	}
	id, err = r.Database.InsertUser(newUser)
	if err != nil {
		log.Println(err)
		RedirectWithError(r, req, res, "Cannot insert user", "/login")
		return
	}

	r.Config.Session.Put(req.Context(), "user_id", id)
	r.Config.Session.Put(req.Context(), "flash", "User registered successfully!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
