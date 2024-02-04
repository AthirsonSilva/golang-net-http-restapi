package repository

import (
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func (r *postgresRepository) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	query := `
						INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id,
																						created_at, updated_at, restriction_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
					`

	_, err := r.DB.SQL.Exec(
		query,
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
