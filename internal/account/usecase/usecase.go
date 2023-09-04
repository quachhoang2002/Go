package usecase

import (
	"context"
	"errors"

	"github.com/quachhoang2002/Go/internal/account/repository"
	"github.com/quachhoang2002/Go/internal/domain"
)

type CreateInput struct {
	Owner    string
	Balance  int32
	Currency string
}

// Create creates a new todo
func (uc implUsecase) Create(ctx context.Context, input CreateInput) error {
	account := domain.Account{
		Owner:    input.Owner,
		Balance:  input.Balance,
		Currency: input.Currency,
	}

	if err := uc.repo.Create(ctx, account); err != nil {
		return err
	}

	return nil
}

// All returns all todos
func (uc implUsecase) All(ctx context.Context) ([]domain.Account, error) {
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
