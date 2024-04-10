-- name: CreateCompanies :one
INSERT INTO companies (
        name,
        document,
        address_id
    )
VALUES ($1, $2, $3)
RETURNING id;
-- name: AssociateUserCompany :one
INSERT INTO company_users (company_id, user_id)
VALUES ($1, $2)
RETURNING *;