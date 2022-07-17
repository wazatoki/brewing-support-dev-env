package repositories

import (
	"brewing_support/app/infrastructures/postgresql"

	"github.com/jmoiron/sqlx"
)

func createDB() *postgresql.Postgresql {
	return postgresql.NewPostgresql()
}

type db interface {
	WithDbContext(fn func(db *sqlx.DB) error) error
}
