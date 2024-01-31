package repository

import (
	"context"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

type postgresRepository struct {
	Config *config.AppConfig
	DB     *database.Database
}

func (r *postgresRepository) InsertReservation(reservation models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var reservationID int

	query := `
						insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
						values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						returning id
					`

	err := r.DB.SQL.QueryRowContext(
		ctx, query,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomID,
		time.Now(),
		time.Now()).Scan(&reservationID)
	if err != nil {
		return 0, err
	}

	return reservationID, nil
}

func (r *postgresRepository) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
						insert into room_restrictions (start_date, end_date, room_id, reservation_id,
																						created_at, updated_at, restriction_id)
						values ($1, $2, $3, $4, $5, $6, $7)
					`

	_, err := r.DB.SQL.ExecContext(
		ctx, query,
		roomRestriction.StartDate,
		roomRestriction.EndDate,
		roomRestriction.RoomID,
		roomRestriction.ReservationID,
		time.Now(),
		time.Now(),
		roomRestriction.RestrictionID,
	)
	if err != nil {
		return err
	}

	return nil
}
