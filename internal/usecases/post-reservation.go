package usecases

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the PostReservation page
func (repo *Repository) PostReservation(res http.ResponseWriter, req *http.Request) {
	reservation, ok := repo.Config.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("Cannot get reservation from session")
		RedirectWithError(repo, req, res, "Cannot get reservation from session", "/")
		return
	}

	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		RedirectWithError(repo, req, res, "Could not parse form", "/")
		return
	}

	raw_start_date := req.Form.Get("start_date")
	raw_end_date := req.Form.Get("end_date")
	layout := "2006-01-02"

	log.Printf("Start date => %s", raw_start_date)
	log.Printf("End date => %s", raw_end_date)

	parsed_start_date, err := time.Parse(layout, raw_start_date)
	if err != nil {
		log.Println(err)
		RedirectWithError(repo, req, res, "Invalid start date", "/")
		return
	}

	parsed_end_date, err := time.Parse(layout, raw_end_date)
	if err != nil {
		log.Println(err)
		RedirectWithError(repo, req, res, "Invalid end date", "/")
		return
	}

	form := forms.New(req.PostForm)
	form.Required("first_name", "last_name", "email", "phone", "start_date", "end_date")
	form.IsEmail("email")
	form.MinLength(2, req, "first_name", "last_name")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(res, req, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	roomID, err := strconv.Atoi(req.Form.Get("room_id"))
	if err != nil {
		log.Println(err)
		RedirectWithError(repo, req, res, "Could not get room id", "/")
		return
	}

	log.Printf("[PostReservation] Getting user id from session")
	userID, err := strconv.Atoi(req.Form.Get("user_id"))
	if err != nil || userID == 0 {
		log.Println("Could not get user id from session")
		RedirectWithError(repo, req, res, "Could not get user id from session", "/login")
		return
	}

	user, err := Repo.Database.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		RedirectWithError(repo, req, res, "Could not get user from session", "/login")
		return
	}

	reservation.RoomID = roomID
	reservation.UserID = userID
	reservation.StartDate = parsed_start_date
	reservation.EndDate = parsed_end_date
	reservation.FirstName = user.FirstName
	reservation.LastName = user.LastName
	reservation.Email = user.Email
	reservation.Phone = user.Phone

	log.Printf(
		"[PostReservation] Reservation => %+v",
		reservation,
	)
	var reservationId int
	reservationId, err = Repo.Database.InsertReservation(reservation)
	if err != nil {
		log.Println(err)
		RedirectWithError(repo, req, res, "Could not make reservation", "/")
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

	log.Printf(
		"[PostReservation] Restriction => %+v",
		restriction,
	)
	err = Repo.Database.InsertRoomRestriction(restriction)
	if err != nil {
		RedirectWithError(repo, req, res, "Could not make reservation", "/")
		return
	}

	log.Println("[PostReservation] Reservation created successfully")

	repo.Config.Session.Put(req.Context(), "reservation", reservation)
	// http.Redirect(res, req, "/reservations/reservation-summary", http.StatusSeeOther)
	repo.ReservationSummary(res, req)
}
