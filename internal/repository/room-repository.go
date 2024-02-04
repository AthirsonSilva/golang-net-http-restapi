package repository

import (
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func (r *postgresRepository) SearchAvailabilityByDateAndRoom(
	start time.Time,
	end time.Time,
	roomID int,
) (bool, error) {
	query := `
						SELECT count(id)
						FROM room_restrictions
						WHERE $1 < end_date 
							AND $2 > start_date
							AND room_id = $3
					`

	var numRows int
	row := r.DB.SQL.QueryRow(query, start, end, roomID)
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
	rows, err := r.DB.SQL.Query(query, start, end)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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
	var room models.Room

	query := `
						SELECT id, room_name, created_at, updated_at
						FROM rooms
						WHERE id = $1
					`

	row := r.DB.SQL.QueryRow(query, roomID)
	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return room, err
	}

	return room, nil
}
