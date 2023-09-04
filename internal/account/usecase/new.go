package usecase

import (
	"context"

	"github.com/quachhoang2002/Go/internal/account/repository"
	"github.com/quachhoang2002/Go/internal/domain"
	"github.com/quachhoang2002/Go/pkg/log"
)

// Usecase is the interface for todo usecase
//
//go:generate mockery --name=Usecase
type Usecase interface {
	// Create creates a new todo
	Create(ctx context.Context, int CreateInput) error
	// All returns all todos
	All(ctx context.Context) ([]domain.Account, error)
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
