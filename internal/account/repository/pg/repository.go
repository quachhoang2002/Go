package pg

import (
	"context"

	"github.com/quachhoang2002/Go/internal/domain"
)

// Create creates a new todo
func (repo implRepository) Create(ctx context.Context, account domain.Account) error {
	createParams := domain.CreateAccountParams{
		Owner:    account.Owner,
		Balance:  account.Balance,
		Currency: account.Currency,
	}

	queries := domain.New(repo.db)
	_, err := queries.CreateAccount(ctx, createParams)
	if err != nil {
		return err
	}

	return nil
}

// All returns all todos
func (repo implRepository) All(ctx context.Context) ([]domain.Account, error) {
	queries := domain.New(repo.db)
	accounts, err := queries.GetListAccount(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// Delete deletes a todo
func (repo implRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
