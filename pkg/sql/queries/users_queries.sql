-- name: CreatePreRegisterUser :one
INSERT INTO users (
    id, 
    name, 
    email, 
    status, 
    register_token, 
    token_expires_at, 
    created_at, 
    role_id, 
    tenant_id
) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, email;

-- name: CreateUserPhones :one
INSERT INTO user_phones (
    id, 
    type, 
    number, 
    is_primary, 
    created_at,
    updated_at
) VALUES($1, $2, $3, $4,$5,$6) RETURNING id, number, type;

-- name: CompleteRegisterUser :one
UPDATE users
SET document_type = $1,
    document_number = $2,
    password = $3,
    status = $4,
    image_id = $5,
    updated_at = $6
WHERE register_token = $7
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
    document_type = $3,
    document_number = $4
WHERE id = $5;

-- name: GetTokenPreRegister :one
SELECT register_token,
    token_expires_at
FROM users
WHERE register_token = $1;

-- name: CreateTenant :one 
INSERT INTO tenant (id, name, tax_code, type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUsersWithEmail :many
SELECT email
FROM users;

-- name: GetUsersContact :many
SELECT email
FROM contact_info;

-- name: GetUserContactEmail :one
SELECT email, name
FROM contact_info
WHERE email = $1;