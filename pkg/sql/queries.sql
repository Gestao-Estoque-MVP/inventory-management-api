-- name: CreateUser :one
INSERT INTO users (name, lastname, email, phone, document_type, document_number, password ) 
    VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateAddress :one
INSERT INTO address (user_id, address, number, street, city, state, postal_code, country) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: DeleteAddress :one
DELETE FROM address WHERE id = $1 RETURNING *;

-- name: GetAddress :one
SELECT * FROM address WHERE id = $1;

-- name: ListAddresses :many
SELECT * FROM address;