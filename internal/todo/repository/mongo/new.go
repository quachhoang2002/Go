package mongo

import (
	"gitlab.com/gma-vietnam/tanca-event/internal/todo/repository"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
)

type implRepository struct {
	l  log.Logger
	db mongo.Database
}

func NewRepository(l log.Logger, db mongo.Database) repository.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
