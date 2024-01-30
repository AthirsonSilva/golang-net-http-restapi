package repository

import (
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
)

type DatabaseRepository interface {
	FindAll() bool
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
