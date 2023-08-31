-- name: CreatePreRegisterUser :one
INSERT INTO users (
        id,
        name,
        email,
        status,
        role_id,
        tenant_id,
        register_token,
        token_expires_at,
        created_at
    )
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id,
    name,
    email;
-- name: CompleteRegisterUser :one
UPDATE users
SET phone = $1,
    document_type = $2,
    document_number = $3,
    password = $4,
    updated_at = $5
WHERE register_token = $6
RETURNING id,
    name,
    email;
-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = $1
RETURNING id,
    name,
    email;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;
-- name: GetEmail :one
SELECT id,
    name,
    email,
    password,
    role_id
FROM users
WHERE email = $1;
-- name: GetUserRegisterToken :one
SELECT *
FROM users
WHERE register_token = $1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id;
-- name: UpdateUser :exec
UPDATE users
SET name = $1,
    email = $2,
    phone = $3,
    document_type = $4,
    document_number = $5
WHERE id = $6;
-- name: GetTokenPreRegister :one
SELECT register_token,
    token_expires_at
FROM users
WHERE register_token = $1;
-- name: CreateTenant :one 
INSERT INTO tenant (id, name)
VALUES ($1, $2)
RETURNING *;
-- name: GetUsersWithEmail :many
SELECT email
FROM users;
-- name: GetUsersContact :many
SELECT *
FROM contact_info;
-- name: GetUserContactEmail :one
SELECT email, name
FROM contact_info
WHERE email = $1;