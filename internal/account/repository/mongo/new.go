package mongo

import (
	"github.com/quachhoang2002/Go/internal/todo/repository"
	"github.com/quachhoang2002/Go/pkg/log"
	"github.com/quachhoang2002/Go/pkg/mongo"
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
