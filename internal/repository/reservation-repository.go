package repository

import (
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func (r *postgresRepository) InsertReservation(reservation models.Reservation) (int, error) {
	var reservationID int

	query := `
						INSERT INTO reservations (user_id, start_date, end_date, room_id, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id
					`

	err := r.DB.SQL.QueryRow(
		query,
		reservation.UserID,
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

func (r *postgresRepository) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	query := `
		SELECT 
			re.id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone,
			re.start_date,
			re.end_date,
			re.room_id,
			ro.room_name,
			re.processed
		FROM reservations re
		LEFT JOIN rooms ro ON (ro.id = re.room_id)
		INNER JOIN users u ON (u.id = re.user_id)
		ORDER BY re.start_date ASC
	`

	rows, err := r.DB.SQL.Query(query)
	if err != nil {
		return reservations, err
	}

	defer rows.Close()

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Room.Name,
			&i.Processed,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, i)
	}

	return reservations, nil
}

func (r *postgresRepository) GetAllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	query := `
		SELECT 
			re.id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone,
			re.start_date,
			re.end_date,
			re.room_id,
			ro.room_name,
			re.processed
		FROM reservations re
		LEFT JOIN rooms ro ON (ro.id = re.room_id)
		INNER JOIN users u ON (u.id = re.user_id)
		WHERE re.processed = 0
		ORDER BY re.start_date asc
	`

	rows, err := r.DB.SQL.Query(query)
	if err != nil {
		return reservations, err
	}

	defer rows.Close()

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Room.Name,
			&i.Processed,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, i)
	}

	return reservations, nil
}

func (r *postgresRepository) GetReservationByID(id int) (models.Reservation, error) {
	var reservation models.Reservation

	query := `
		SELECT 
			re.id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone,
			re.start_date,
			re.end_date,
			re.room_id,
			ro.room_name,
			re.processed
		FROM reservations re
		LEFT JOIN rooms ro ON (ro.id = re.room_id)
		INNER JOIN users u on u.id = re.user_id
		WHERE re.id = $1
		ORDER BY re.start_date asc
	`

	row := r.DB.SQL.QueryRow(query, id)

	err := row.Scan(
		&reservation.ID,
		&reservation.FirstName,
		&reservation.LastName,
		&reservation.Email,
		&reservation.Phone,
		&reservation.StartDate,
		&reservation.EndDate,
		&reservation.RoomID,
		&reservation.Room.Name,
		&reservation.Processed,
	)
	if err != nil {
		return reservation, err
	}

	return reservation, nil
}

func (r *postgresRepository) DeleteReservationByID(id int) error {
	query := `
		DELETE FROM reservations
		WHERE id = $1
	`

	_, err := r.DB.SQL.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) UpdateReservation(reservation models.Reservation) error {
	query := `
		UPDATE reservations
		SET 
			start_date = $5,
			end_date = $6,
			processed = $7
		WHERE id = $8
		`

	_, err := r.DB.SQL.Exec(
		query,
		reservation.StartDate,
		reservation.EndDate,
		reservation.Processed,
		reservation.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
