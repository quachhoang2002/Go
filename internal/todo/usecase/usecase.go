package usecase

import (
	"context"
	"errors"

	"github.com/quachhoang2002/Go/internal/model"
	"github.com/quachhoang2002/Go/internal/todo/repository"
)

type CreateInput struct {
	Name        string
	Description string
}

// Create creates a new todo
func (uc implUsecase) Create(ctx context.Context, input CreateInput) error {
	m := model.Todo{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := uc.repo.Create(ctx, m); err != nil {
		uc.l.Errorf(ctx, "todo.usecase.Create.repo.Create(ctx, %+v): %s", m, err)
		return err
	}

	return nil
}

// All returns all todos
func (uc implUsecase) All(ctx context.Context) ([]model.Todo, error) {
	todos, err := uc.repo.All(ctx)
	if err != nil {
		uc.l.Errorf(ctx, "todo.usecase.All.repo.All(ctx): %s", err)
		return nil, err
	}

	return todos, nil
}

// Delete deletes a todo
func (uc implUsecase) Delete(ctx context.Context, id int) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			uc.l.Warnf(ctx, "todo.usecase.Delete.repo.Delete(ctx, %d): %s", id, err)
			return ErrNotFound
		}

		uc.l.Errorf(ctx, "todo.usecase.Delete.repo.Delete(ctx, %d): %s", id, err)
		return err
	}

	return nil
}
