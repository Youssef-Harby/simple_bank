-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: ListTransfersBetweenAccounts :many
SELECT * FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2
ORDER BY created_at DESC
OFFSET $3
LIMIT $4;

-- name: UpdateTransfer :one
UPDATE transfers 
SET amount = $2 
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;