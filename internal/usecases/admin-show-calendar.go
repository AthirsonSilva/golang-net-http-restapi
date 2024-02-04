package usecases

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
)

func (r *Repository) AdminReservationsCalendar(res http.ResponseWriter, req *http.Request) {
	now := time.Now()

	if req.URL.Query().Get("y") != "" && req.URL.Query().Get("m") != "" {
		year, err := strconv.Atoi(req.URL.Query().Get("y"))
		if err != nil {
			helpers.ServerError(res, err)
		}

		month, err := strconv.Atoi(req.URL.Query().Get("m"))
		if err != nil {
			helpers.ServerError(res, err)
		}

		now = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}

	next := now.AddDate(0, 1, 0)
	last := now.AddDate(0, -1, 0)

	nextMonth := next.Format("01")
	nextMonthYear := next.Format("2006")

	lastMonth := last.Format("01")
	lastMonthYear := last.Format("2006")

	dateMap := make(map[string]string)
	dateMap["next_month"] = nextMonth
	dateMap["next_month_year"] = nextMonthYear
	dateMap["last_month"] = lastMonth
	dateMap["last_month_year"] = lastMonthYear

	dateMap["this_month"] = now.Format("01")
	dateMap["this_year"] = now.Format("2006")

	render.RenderTemplate(res, req, "admin-reservations-calendar.page.tmpl", &models.TemplateData{
		DateMap: dateMap,
	})
}
