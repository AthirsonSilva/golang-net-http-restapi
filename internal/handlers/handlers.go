package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

var Repo *Repository

// System-wide configuration struct
type Repository struct {
	Config *config.AppConfig
}

// Creates a new Repo (Application Config) instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		Config: a,
	}
}

// Creates a new Handlers instance
func NewHandlers(r *Repository) {
	Repo = r
}

// Responsible for the Home page
func (repo *Repository) Home(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "home.page.tmpl", &models.TemplateData{})
}

// Responsible for the About page
func (repo *Repository) About(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "about.page.tmpl", &models.TemplateData{})
}

// Responsible for the MakeReservation page
func (repo *Repository) MakeReservation(responseWriter http.ResponseWriter, request *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(responseWriter, request, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// Responsible for the PostReservation page
func (repo *Repository) PostReservation(responseWriter http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	// Get the post data from the form
	reservation := models.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Phone:     request.Form.Get("phone"),
		Email:     request.Form.Get("email"),
	}

	// Apply validation for every form field
	form := forms.New(request.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.IsEmail("email")
	for _, field := range []string{"first_name", "last_name"} {
		form.MinLength(field, 2, request)
	}

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(responseWriter, request, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	repo.Config.Session.Put(request.Context(), "reservation", reservation)
	http.Redirect(responseWriter, request, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummary(responseWriter http.ResponseWriter, request *http.Request) {
	reservation, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		err := errors.New("cannot get reservation from session")
		helpers.ServerError(responseWriter, err)
		repo.Config.Session.Put(request.Context(), "error", "Can't get reservation from session")
		http.Redirect(responseWriter, request, "/", http.StatusTemporaryRedirect)
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation
	repo.Config.Session.Remove(request.Context(), "reservation")

	// Render the Reservation Summary page
	render.RenderTemplate(responseWriter, request, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Responsible for rendering the Availability page
func (repo *Repository) Availability(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "search-availability.page.tmpl", &models.TemplateData{})
}

// Responsible for receiving the data from the Availability page
func (repo *Repository) PostAvailability(responseWriter http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	responseWriter.Write([]byte(fmt.Sprintf(
		"start date is %s and end date is %s", start, end,
	)))
}

// Responsible for rendering the Availability JSON page
func (repo *Repository) PostAvailabilityJSON(responseWriter http.ResponseWriter, request *http.Request) {
	response := models.JsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(response, "", "     ")
	if err != nil {
		helpers.ServerError(responseWriter, err)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(out)
}

func (repo *Repository) Contact(responseWriter http.ResponseWriter, request *http.Request) {
	// Responsible for rendering the Contact page
	render.RenderTemplate(responseWriter, request, "contact.page.tmpl", &models.TemplateData{})
}

// Responsible for rendering the Reservation Summary page
func (repo *Repository) Majors(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "majors.page.tmpl", &models.TemplateData{})
}

// Responsible for rendering the Reservation Summary page
func (repo *Repository) Generals(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "generals.page.tmpl", &models.TemplateData{})
}
