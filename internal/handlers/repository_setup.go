package handlers

import (
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/repository"
)

var Repo *Repository

// System-wide configuration struct
type Repository struct {
	Config   *config.AppConfig
	Database repository.DatabaseRepository
}

// Creates a new Repo (Application Config) instance
func NewRepo(ac *config.AppConfig, db *database.Database) *Repository {
	return &Repository{
		Config:   ac,
		Database: repository.NewPostgresRepository(ac, db),
	}
}

// Creates a new Handlers instance
func NewHandlers(r *Repository) {
	Repo = r
}
