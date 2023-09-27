-- name: GetRole :one
SELECT * FROM roles WHERE name = $1;

-- name: GetRoleByID :one
SELECT * from roles WHERE id = $1;


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

-- name: CreateUsersRoles :one

INSERT INTO users_roles (user_id, role_id) 
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

-- name: GetRoleUser :one
SELECT r.id, r.name
FROM users u
JOIN users_roles ur ON u.id = ur.user_id
JOIN roles r ON ur.role_id = r.id
WHERE u.id = $1; 
