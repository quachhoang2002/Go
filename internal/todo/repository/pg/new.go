package pg

import (
	"database/sql"

	"github.com/quachhoang2002/Go/internal/todo/repository"
	"github.com/quachhoang2002/Go/pkg/log"
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
