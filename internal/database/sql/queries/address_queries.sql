-- name: CreateAddress :one
INSERT INTO address (
        id,
        address,
        number,
        street,
        city,
        state,
        postal_code,
        country,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;