-- name: CreatePreRegisterUser :one
INSERT INTO users (id, name, email, status, role_id, tenant_id, register_token, token_expires_at, created_at) 
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, email;

-- name: CompleteRegisterUser :one
UPDATE users SET phone = $1, document_type = $2, document_number = $3, password = $4, avatar = $5 WHERE register_token = $6 RETURNING id, name, email;

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = $1 RETURNING id, name, email;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserRegisterToken :one
SELECT * FROM users WHERE register_token = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateUser :one
UPDATE users SET name = $1, email = $2, phone = $3, document_type = $4, document_number = $5 WHERE id = $6 RETURNING id, name, email;


 
-- name: CreateAddress :one
INSERT INTO address (user_id, address, number, street, city, state, postal_code, country) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: DeleteAddress :execresult
DELETE FROM address WHERE id = $1 RETURNING *;

-- name: UpdateAddress :one
UPDATE address SET user_id = $1, address = $2, number = $3, street = $4, city = $5, state = $6, postal_code = $7, country = $8 WHERE id = $9 RETURNING *;

-- name: GetAddress :one
SELECT * FROM address WHERE id = $1;

-- name: ListAddresses :many
SELECT * FROM address;

-- name: GetRole :one
SELECT * FROM roles WHERE name = $1;


-- name: CreateContactInfo :one
INSERT INTO contact_info (id, name, email, phone, created_at) 
    VALUES ($1, $2, $3, $4, $5) RETURNING *;


-- name: CreatePermissions :one 
INSERT INTO permissions (id, name, description) 
    VALUES ($1, $2, $3) RETURNING *;

-- name: CreateRole :one
INSERT INTO roles (id, name, description) 
    VALUES ($1, $2, $3) RETURNING *;

-- name: CreateRolePermissions :one
INSERT INTO roles_permissions (role_id, permission_id) 
    VALUES ($1, $2) RETURNING *;

-- name: GetRolesPermissions :many
SELECT
    u.id AS user_id,
    r.name AS role_name,
    p.name AS permission_name
FROM 
    users AS u
INNER JOIN 
    roles AS r ON u.role_id = r.id
INNER JOIN 
    roles_permissions AS rp ON r.id = rp.role_id
INNER JOIN
    permissions AS p ON rp.permission_id = p.id
WHERE 
    u.id = $1;

-- name: CreateTenant :one 
INSERT INTO tenant (id, name) 
    VALUES ($1, $2) RETURNING *;

-- name: CreateUsersPermissions :one
INSERT INTO users_permissions (user_id, permission_id) 
    VALUES ($1, $2) RETURNING *;

-- name: GetUsersPermissions :many
SELECT
    u.id AS user_id,
    p.name AS permission_name
FROM
    users AS u
INNER JOIN
    users_permissions AS up ON u.id = up.user_id
INNER JOIN
    permissions AS p ON up.permission_id = p.id
WHERE
    u.id = $1;
    