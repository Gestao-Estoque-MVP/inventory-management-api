-- name: CreateAddress :one
INSERT INTO address (user_id, address, number, street, city, state, postal_code, country, created_at) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: DeleteAddress :execresult
DELETE FROM address WHERE user_id = $1 RETURNING *;

-- name: UpdateAddress :one
UPDATE address SET user_id = $1, address = $2, number = $3, street = $4, city = $5, state = $6, postal_code = $7, country = $8 WHERE id = $9 RETURNING *;

-- name: GetAddressByID :one
SELECT * FROM address WHERE user_id = $1;

-- name: ListAddresses :many
SELECT * FROM address;