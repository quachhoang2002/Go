package repository

import (
	"context"

	"github.com/quachhoang2002/Go/internal/domain"
)

// Repository is the interface for todo repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create creates a new todo
	Create(ctx context.Context, account domain.Account) error
	// All returns all todos
	All(ctx context.Context) ([]domain.Account, error)
	// Delete deletes a todo
	Delete(ctx context.Context, id int) error
}
