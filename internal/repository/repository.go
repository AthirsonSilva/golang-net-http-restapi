package repository

import (
	"context"
	"errors"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"golang.org/x/crypto/bcrypt"
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
						INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						RETURNING id
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
						INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id,
																						created_at, updated_at, restriction_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
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

func (r *postgresRepository) SearchAvailabilityByDateAndRoom(
	start time.Time,
	end time.Time,
	roomID int,
) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
						SELECT count(id)
						FROM room_restrictions
						WHERE $1 < end_date 
							AND $2 > start_date
							AND room_id = $3
					`

	var numRows int
	row := r.DB.SQL.QueryRowContext(ctx, query, start, end, roomID)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

func (r *postgresRepository) SearchAvailabilityByDateForAllRooms(
	start time.Time,
	end time.Time,
) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
						SELECT r.id, r.room_name
						FROM rooms r
						WHERE r.id IN (
							SELECT rr.room_id
							FROM room_restrictions rr
							WHERE $1 < rr.end_date
							AND $2 > rr.start_date
						)
					`

	var rooms []models.Room
	rows, err := r.DB.SQL.QueryContext(ctx, query, start, end)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *postgresRepository) GetRoomByID(roomID int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var room models.Room

	query := `
						SELECT id, room_name, created_at, updated_at
						FROM rooms
						WHERE id = $1
					`

	row := r.DB.SQL.QueryRowContext(ctx, query, roomID)
	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *postgresRepository) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	query := `
						SELECT 
							id, 
							first_name, 
							last_name, 
							email, 
							password, 
							access_level, 
							created_at, 
							updated_at
						FROM users 
						WHERE id = $1
					`

	row := r.DB.SQL.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *postgresRepository) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE users 
		SET first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5
		WHERE id = $6
	`

	_, err := r.DB.SQL.ExecContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.AccessLevel,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) Authenticate(email string, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	query := `
		SELECT 
			id,
			password
		FROM users
		WHERE email = $1
	`

	row := r.DB.SQL.QueryRowContext(ctx, query, email)

	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}
