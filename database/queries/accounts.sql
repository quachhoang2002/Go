-- name: GetListAccount :many
SELECT
    *
FROM
    accounts;

-- name: CreateAccount :one
INSERT INTO
    accounts (owner, balance, currency)
VALUES
    ($1, $2, $3) RETURNING *;