package usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/forms"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func (r *Repository) AdminUpdateReservation(res http.ResponseWriter, req *http.Request) {
	log.Printf("[AdminUpdateReservation] Updating reservation...")
	var reservation models.Reservation

	err := req.ParseForm()
	if err != nil {
		helpers.ServerError(res, err)
	}

	id, err := strconv.Atoi(req.Form.Get("id"))
	if err != nil {
		helpers.ServerError(res, err)
	}

	firstName := req.Form.Get("first_name")
	lastName := req.Form.Get("last_name")
	email := req.Form.Get("email")
	phone := req.Form.Get("phone")
	startDate := req.Form.Get("start_date")
	endDate := req.Form.Get("end_date")

	form := forms.New(req.Form)
	form.Required("firstName", "lastName", "email", "phone", "start_date", "end_date")
	form.MinLength(2, req, "firstName", "lastName", "email", "phone", "start_date", "end_date")
	form.IsEmail("email")

	parsedStartDate := helpers.ConvertDateFromString(startDate, res)
	parsedEndDate := helpers.ConvertDateFromString(endDate, res)

	newReservation := models.Reservation{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		StartDate: parsedStartDate,
		EndDate:   parsedEndDate,
	}

	err = r.Database.UpdateReservation(newReservation)

	data := make(map[string]interface{})
	data["reservation"] = reservation

	http.Redirect(res, req, "/admin/reservations/all", http.StatusSeeOther)
}
