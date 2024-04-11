-- name: CreateCompanies :one
INSERT INTO companies (
        name,
        document,
        address_id,
        is_admin
    )
VALUES ($1, $2, $3, $4)
RETURNING id;
-- name: AssociateUserCompany :one
INSERT INTO company_users (company_id, user_id)
VALUES ($1, $2)
RETURNING *;