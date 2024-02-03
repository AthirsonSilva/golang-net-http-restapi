package usecases

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

// Responsible for the PostReservation page
func (repo *Repository) PostReservation(res http.ResponseWriter, req *http.Request) {
	reservation, ok := repo.Config.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(res, errors.New("cannot get reservation from session"))
		return
	}

	err := req.ParseForm()
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	err = req.ParseForm()
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	raw_start_date := req.Form.Get("start_date")
	raw_end_date := req.Form.Get("end_date")
	layout := "2006-01-02"

	parsed_start_date, err := time.Parse(layout, raw_start_date)
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	parsed_end_date, err := time.Parse(layout, raw_end_date)
	if err != nil {
		helpers.ServerError(res, err)
		return
	}

	form := forms.New(req.PostForm)
	form.Required("first_name", "last_name", "email", "phone", "start_date", "end_date")
	form.IsEmail("email")
	for _, field := range []string{"first_name", "last_name"} {
		form.MinLength(field, 2, req)
	}

	if err != nil {
		helpers.ServerError(res, err)
		return
	}

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
		helpers.ServerError(res, err)
		return
	}

	reservation.FirstName = req.Form.Get("first_name")
	reservation.LastName = req.Form.Get("last_name")
	reservation.Email = req.Form.Get("email")
	reservation.Phone = req.Form.Get("phone")
	reservation.RoomID = roomID
	reservation.StartDate = parsed_start_date
	reservation.EndDate = parsed_end_date

	var reservationId int
	reservationId, err = Repo.Database.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(res, err)
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
		helpers.ServerError(res, err)
		return
	}

	repo.Config.Session.Put(req.Context(), "reservation", res)
	http.Redirect(res, req, "/reservation-summary", http.StatusSeeOther)
}
