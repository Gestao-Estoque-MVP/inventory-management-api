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