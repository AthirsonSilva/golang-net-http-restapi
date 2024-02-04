package repository

import (
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

type DatabaseRepository interface {
	// Reservations methods
	InsertReservation(reservation models.Reservation) (int, error)
	GetReservationByID(id int) (models.Reservation, error)
	GetAllReservations() ([]models.Reservation, error)
	GetAllNewReservations() ([]models.Reservation, error)
	DeleteReservationByID(id int) error
	UpdateReservation(reservation models.Reservation) error

	// RoomRestriction methods
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error

	// Rooms methods
	SearchAvailabilityByDateAndRoom(start time.Time, end time.Time, roomID int) (bool, error)
	SearchAvailabilityByDateForAllRooms(start time.Time, end time.Time) ([]models.Room, error)
	GetRoomByID(roomID int) (models.Room, error)

	// User methods
	GetUserByID(id int) (models.User, error)
	GetUserByEmailAndPassword(email string, testPassword string) (int, string, error)
	UpdateUser(user models.User) error
}

type postgresRepository struct {
	Config *config.AppConfig
	DB     *database.Database
}

func NewPostgresRepository(config *config.AppConfig, db *database.Database) DatabaseRepository {
	return &postgresRepository{
		Config: config,
		DB:     db,
	}
}
