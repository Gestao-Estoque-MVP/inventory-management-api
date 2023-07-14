-- name: CreateUser :exec
INSERT INTO users (name, lastname, email, document_type, document_number, address_id ) 
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1 RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateAddress :exec
INSERT INTO address (id, street, city, state, zip_code) 
    VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address WHERE id = $1 RETURNING *;

-- name: GetAddress :one
SELECT * FROM address WHERE id = $1;

-- name: ListAddresses :many
SELECT * FROM address;