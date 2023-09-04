-- name: GetListTransfer :one
SELECT
    *
FROM
    transfer;

-- name: GetTransfer :one
SELECT
    *
FROM 
    transfer
WHERE
    id = $1 LIMIT 1;

-- name: CreateTransfer :one
INSERT INTO
    transfer (from_account_id, to_account_id, amount)
VALUES
    ($1, $2, $3) RETURNING *;

