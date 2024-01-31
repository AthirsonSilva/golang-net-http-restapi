package repository

import (
	"context"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func (r *postgresRepository) FindAllUsers() bool {
	return true
}

func (r *postgresRepository) InsertReservation(reservation models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
						insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
						values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						returning id
					`

	_, err := r.DB.SQL.ExecContext(
		ctx, query,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomID,
		time.Now(),
		time.Now())
	if err != nil {
		return err
	}

	return nil
}
