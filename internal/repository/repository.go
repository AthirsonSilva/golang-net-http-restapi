package repository

import (
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

type DatabaseRepository interface {
	FindAllUsers() bool

	InsertReservation(reservation models.Reservation) error
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
