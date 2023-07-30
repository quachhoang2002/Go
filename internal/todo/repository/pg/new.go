package pg

import (
	"database/sql"

	"gitlab.com/gma-vietnam/tanca-event/internal/todo/repository"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type implRepository struct {
	l  log.Logger
	db *sql.DB
}

func NewRepository(l log.Logger, db *sql.DB) repository.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
