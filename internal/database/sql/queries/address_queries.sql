-- name: CreateAddress :one
INSERT INTO address (
        address,
        number,
        street,
        city,
        state,
        postal_code,
        country
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;