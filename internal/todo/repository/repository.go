package repository

import (
	"context"

	"github.com/quachhoang2002/Go/internal/model"
)

// Repository is the interface for todo repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create creates a new todo
	Create(ctx context.Context, todo model.Todo) error
	// All returns all todos
	All(ctx context.Context) ([]model.Todo, error)
	// Delete deletes a todo
	Delete(ctx context.Context, id int) error
}
