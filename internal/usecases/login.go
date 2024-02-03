package usecases

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
)

func (r *Repository) Login(responseWriter http.ResponseWriter, request *http.Request) {
	_ = r.Config.Session.RenewToken(request.Context())

	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := request.Form.Get("email")
	password := request.Form.Get("password")

	form := forms.New(request.Form)
	form.Required("email", "password")
	if !form.Valid() {
		http.Redirect(responseWriter, request, "/", http.StatusSeeOther)
		return
	}

	id, _, err := r.Database.Authenticate(email, password)
	if err != nil {
		log.Println(err)
		r.Config.Session.Put(request.Context(), "error", "Invalid login credentials")
		http.Redirect(responseWriter, request, "/login", http.StatusSeeOther)
	}

	r.Config.Session.Put(request.Context(), "user_id", id)
	r.Config.Session.Put(request.Context(), "flash", "Logged in successfully!")
	http.Redirect(responseWriter, request, "/", http.StatusSeeOther)
}
