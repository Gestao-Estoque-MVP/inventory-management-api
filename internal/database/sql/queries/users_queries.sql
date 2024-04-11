-- name: CreateUser :one
INSERT INTO users (
        name,
        email,
        document,
        password,
        mobile_phone,
        active,
        register_token,
        token_expires_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;
-- name: GetUserByEmail :one
SELECT u.id,
    email,
    password,
    r.name AS role_name
FROM users u
    JOIN users_roles ur ON u.id = ur.user_id
    JOIN roles r ON ur.role_id = r.id
WHERE email = $1;