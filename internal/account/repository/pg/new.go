package pg

import (
	"database/sql"

	"github.com/quachhoang2002/Go/internal/account/repository"
	"github.com/quachhoang2002/Go/pkg/log"
)

type implRepository struct {
	l  log.Logger
	db *sql.DB
}

var _ repository.Repository = implRepository{}

func NewRepository(l log.Logger, db *sql.DB) repository.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
