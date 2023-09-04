package pg

import (
	"context"

	"github.com/quachhoang2002/Go/internal/domain"
)

// Create creates a new todo
func (repo implRepository) Create(ctx context.Context, account domain.Account) error {
	panic("not implemented")
}

// All returns all todos
func (repo implRepository) All(ctx context.Context) ([]domain.Account, error) {
	panic("not implemented")
}

// Delete deletes a todo
func (repo implRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
