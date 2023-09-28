-- name: CreatePreRegisterUser :one
WITH inserted_user AS (
    INSERT INTO users (
        id, 
        name, 
        email, 
        status, 
        register_token, 
        token_expires_at, 
        created_at, 
        tenant_id
    ) 
    VALUES($1, $2, $3, $4, $5, $6, $7, $8) 
    RETURNING id
) INSERT INTO user_phones (
    id, 
    type, 
    number, 
    is_primary,
    user_id, 
    created_at,
    updated_at
) 
VALUES ($9, $10, $11, $12,(SELECT id FROM inserted_user), $13, $14)
RETURNING (SELECT id FROM inserted_user) AS id;

-- name: CreateUserPhones :one
INSERT INTO user_phones (
    id, 
    type, 
    number, 
    is_primary,
    user_id, 
    created_at,
    updated_at
) VALUES($1, $2, $3, $4,$5,$6,$7) RETURNING id, number, type;

-- name: CompleteRegisterUser :one
WITH inserted_image AS (
    INSERT INTO image (
        id,
        url,
        description,
        created_at
    ) 
    VALUES($7, $8, $9, $10)
    RETURNING id AS image_id
)
UPDATE users
SET 
    document_type = $1,
    document_number = $2,
    password = $3,
    status = $4,
    image_id = (SELECT image_id FROM inserted_image),
    updated_at = $5
WHERE register_token = $6
RETURNING id, name, email;


-- name: CreateCompanyUsers :one
INSERT INTO users (
    id, 
    name, 
    email, 
    status, 
    register_token, 
    token_expires_at, 
    created_at, 
    tenant_id
) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, name, email;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = $1
RETURNING id,
    name,
    email;

-- name: GetUser :one
SELECT users.*, sqlc.embed(address), sqlc.embed(image), sqlc.embed(user_phones)
FROM users
LEFT JOIN address ON address.user_id = users.id
LEFT JOIN image ON image.id = users.image_id
LEFT JOIN user_phones ON user_phones.user_id = users.id
WHERE users.id = $1;

-- name: GetEmail :one
SELECT id,
    name,
    email,
    password
FROM users
WHERE email = $1;

-- name: GetUserRegisterToken :one
SELECT id, name, register_token
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