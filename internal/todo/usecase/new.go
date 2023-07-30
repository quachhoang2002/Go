package usecase

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/model"
	"gitlab.com/gma-vietnam/tanca-event/internal/todo/repository"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

// Usecase is the interface for todo usecase
//
//go:generate mockery --name=Usecase
type Usecase interface {
	// Create creates a new todo
	Create(ctx context.Context, int CreateInput) error
	// All returns all todos
	All(ctx context.Context) ([]model.Todo, error)
	// Delete deletes a todo
	Delete(ctx context.Context, id int) error
}

type implUsecase struct {
	l    log.Logger
	repo repository.Repository
}

func New(l log.Logger, repo repository.Repository) Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
