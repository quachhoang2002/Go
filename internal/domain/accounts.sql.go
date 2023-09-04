// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: accounts.sql

package domain

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO
    accounts (owner, balance, currency)
VALUES
    ($1, $2, $3) RETURNING id, owner, balance, currency, created_at, updated_at
`

type CreateAccountParams struct {
	Owner    string
	Balance  int32
	Currency string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListAccount = `-- name: GetListAccount :many
SELECT
    id, owner, balance, currency, created_at, updated_at
FROM
    accounts
`

func (q *Queries) GetListAccount(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getListAccount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
