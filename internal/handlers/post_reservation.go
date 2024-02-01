package handlers

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
