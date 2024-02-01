package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/repository"
	"github.com/go-chi/chi/v5"
)

var Repo *Repository

// System-wide configuration struct
type Repository struct {
	Config   *config.AppConfig
	Database repository.DatabaseRepository
}

// Creates a new Repo (Application Config) instance
func NewRepo(ac *config.AppConfig, db *database.Database) *Repository {
	return &Repository{
		Config:   ac,
		Database: repository.NewPostgresRepository(ac, db),
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
	res, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(responseWriter, errors.New("cannot get reservation from session"))
		return
	}

	room, err := repo.Database.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	res.Room.RoomName = room.RoomName

	repo.Config.Session.Put(request.Context(), "reservation", res)

	data := make(map[string]interface{})
	data["reservation"] = res

	dateMap := make(map[string]string)
	dateMap["start_date"] = res.StartDate.Format("2006-01-02")
	dateMap["end_date"] = res.EndDate.Format("2006-01-02")

	render.RenderTemplate(responseWriter, request, "make-reservation.page.tmpl", &models.TemplateData{
		Form:    forms.New(nil),
		Data:    data,
		DateMap: dateMap,
	})
}

// Responsible for the PostReservation page
func (repo *Repository) PostReservation(responseWriter http.ResponseWriter, request *http.Request) {
	res, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(responseWriter, errors.New("cannot get reservation from session"))
		return
	}

	err := request.ParseForm()
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	err = request.ParseForm()
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	raw_start_date := request.Form.Get("start_date")
	raw_end_date := request.Form.Get("end_date")
	layout := "2006-01-02"

	parsed_start_date, err := time.Parse(layout, raw_start_date)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	parsed_end_date, err := time.Parse(layout, raw_end_date)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	form := forms.New(request.PostForm)
	form.Required("first_name", "last_name", "email", "phone", "start_date", "end_date")
	form.IsEmail("email")
	for _, field := range []string{"first_name", "last_name"} {
		form.MinLength(field, 2, request)
	}

	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = res

		render.RenderTemplate(responseWriter, request, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	roomID, err := strconv.Atoi(request.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	res.FirstName = request.Form.Get("first_name")
	res.LastName = request.Form.Get("last_name")
	res.Email = request.Form.Get("email")
	res.Phone = request.Form.Get("phone")
	res.RoomID = roomID
	res.StartDate = parsed_start_date
	res.EndDate = parsed_end_date

	var reservationId int
	reservationId, err = Repo.Database.InsertReservation(res)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	defaultRestriction := 1
	restriction := models.RoomRestriction{
		StartDate:     parsed_start_date,
		EndDate:       parsed_end_date,
		RoomID:        roomID,
		ReservationID: reservationId,
		RestrictionID: defaultRestriction,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = Repo.Database.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	repo.Config.Session.Put(request.Context(), "reservation", res)
	http.Redirect(responseWriter, request, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummary(responseWriter http.ResponseWriter, request *http.Request) {
	reservation, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(responseWriter, errors.New("cannot get reservation"))
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation
	repo.Config.Session.Remove(request.Context(), "reservation")

	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	dateMap := make(map[string]string)
	dateMap["start_date"] = startDate
	dateMap["end_date"] = endDate

	// Render the Reservation Summary page
	render.RenderTemplate(responseWriter, request, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:    data,
		DateMap: dateMap,
	})
}

// Responsible for rendering the Availability page
func (repo *Repository) Availability(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, request, "search-availability.page.tmpl", &models.TemplateData{})
}

// Responsible for receiving the data from the Availability page
func (repo *Repository) PostAvailability(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	layout := "2006-01-02"
	parsed_start_date, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	parsed_end_date, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	rooms, err := Repo.Database.SearchAvailabilityByDateForAllRooms(parsed_start_date, parsed_end_date)
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	for _, room := range rooms {
		log.Printf("Found room: %v", room.RoomName)
	}

	if len(rooms) == 0 {
		Repo.Config.Session.Put(request.Context(), "error", "No availability")
		http.Redirect(responseWriter, request, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: parsed_start_date,
		EndDate:   parsed_end_date,
		RoomID:    0,
	}

	repo.Config.Session.Put(request.Context(), "reservation", res)

	render.RenderTemplate(responseWriter, request, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
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

// Responsible for rendering the Choose room page
func (repo *Repository) ChooseRoom(responseWriter http.ResponseWriter, request *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		helpers.ServerError(responseWriter, err)
		return
	}

	res, ok := repo.Config.Session.Get(request.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(responseWriter, err)
		return
	}

	res.RoomID = roomID
	repo.Config.Session.Put(request.Context(), "reservation", res)

	http.Redirect(responseWriter, request, "/make-reservation", http.StatusSeeOther)
}
