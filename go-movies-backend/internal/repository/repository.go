package repository

import (
	"database/sql"

	"github.com/huannguyen2114/Another-Golang-Project/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
