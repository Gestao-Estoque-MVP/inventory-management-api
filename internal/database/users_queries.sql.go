// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users_queries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const completeRegisterUser = `-- name: CompleteRegisterUser :one
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
RETURNING id, name, email
`

type CompleteRegisterUserParams struct {
	DocumentType   pgtype.Text
	DocumentNumber pgtype.Text
	Password       pgtype.Text
	Status         UserStatus
	UpdatedAt      pgtype.Timestamp
	RegisterToken  pgtype.Text
	ID             pgtype.UUID
	Url            string
	Description    pgtype.Text
	CreatedAt      pgtype.Timestamp
}

type CompleteRegisterUserRow struct {
	ID    pgtype.UUID
	Name  pgtype.Text
	Email string
}

func (q *Queries) CompleteRegisterUser(ctx context.Context, arg CompleteRegisterUserParams) (CompleteRegisterUserRow, error) {
	row := q.db.QueryRow(ctx, completeRegisterUser,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.Password,
		arg.Status,
		arg.UpdatedAt,
		arg.RegisterToken,
		arg.ID,
		arg.Url,
		arg.Description,
		arg.CreatedAt,
	)
	var i CompleteRegisterUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createCompanyUsers = `-- name: CreateCompanyUsers :one
INSERT INTO users (
    id, 
    name, 
    email, 
    status, 
    register_token, 
    token_expires_at, 
    created_at, 
    tenant_id
) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, name, email
`

type CreateCompanyUsersParams struct {
	ID             pgtype.UUID
	Name           pgtype.Text
	Email          string
	Status         UserStatus
	RegisterToken  pgtype.Text
	TokenExpiresAt pgtype.Timestamp
	CreatedAt      pgtype.Timestamp
	TenantID       string
}

type CreateCompanyUsersRow struct {
	ID    pgtype.UUID
	Name  pgtype.Text
	Email string
}

func (q *Queries) CreateCompanyUsers(ctx context.Context, arg CreateCompanyUsersParams) (CreateCompanyUsersRow, error) {
	row := q.db.QueryRow(ctx, createCompanyUsers,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Status,
		arg.RegisterToken,
		arg.TokenExpiresAt,
		arg.CreatedAt,
		arg.TenantID,
	)
	var i CreateCompanyUsersRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createPreRegisterUser = `-- name: CreatePreRegisterUser :one
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
RETURNING (SELECT id FROM inserted_user) AS id
`

type CreatePreRegisterUserParams struct {
	ID             pgtype.UUID
	Name           pgtype.Text
	Email          string
	Status         UserStatus
	RegisterToken  pgtype.Text
	TokenExpiresAt pgtype.Timestamp
	CreatedAt      pgtype.Timestamp
	TenantID       string
	ID_2           pgtype.UUID
	Type           TypeNumber
	Number         string
	IsPrimary      bool
	CreatedAt_2    pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

func (q *Queries) CreatePreRegisterUser(ctx context.Context, arg CreatePreRegisterUserParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createPreRegisterUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Status,
		arg.RegisterToken,
		arg.TokenExpiresAt,
		arg.CreatedAt,
		arg.TenantID,
		arg.ID_2,
		arg.Type,
		arg.Number,
		arg.IsPrimary,
		arg.CreatedAt_2,
		arg.UpdatedAt,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenant (id, name, tax_code, type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, tax_code, type, created_at, updated_at
`

type CreateTenantParams struct {
	ID        pgtype.UUID
	Name      pgtype.Text
	TaxCode   pgtype.Text
	Type      NullTenantType
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRow(ctx, createTenant,
		arg.ID,
		arg.Name,
		arg.TaxCode,
		arg.Type,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TaxCode,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserPhones = `-- name: CreateUserPhones :one
INSERT INTO user_phones (
    id, 
    type, 
    number, 
    is_primary,
    user_id, 
    created_at,
    updated_at
) VALUES($1, $2, $3, $4,$5,$6,$7) RETURNING id, number, type
`

type CreateUserPhonesParams struct {
	ID        pgtype.UUID
	Type      TypeNumber
	Number    string
	IsPrimary bool
	UserID    pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type CreateUserPhonesRow struct {
	ID     pgtype.UUID
	Number string
	Type   TypeNumber
}

func (q *Queries) CreateUserPhones(ctx context.Context, arg CreateUserPhonesParams) (CreateUserPhonesRow, error) {
	row := q.db.QueryRow(ctx, createUserPhones,
		arg.ID,
		arg.Type,
		arg.Number,
		arg.IsPrimary,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i CreateUserPhonesRow
	err := row.Scan(&i.ID, &i.Number, &i.Type)
	return i, err
}

const deleteUser = `-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = $1
RETURNING id,
    name,
    email
`

type DeleteUserRow struct {
	ID    pgtype.UUID
	Name  pgtype.Text
	Email string
}

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteUser, id)
}

const getEmail = `-- name: GetEmail :one
SELECT id,
    name,
    email,
    password
FROM users
WHERE email = $1
`

type GetEmailRow struct {
	ID       pgtype.UUID
	Name     pgtype.Text
	Email    string
	Password pgtype.Text
}

func (q *Queries) GetEmail(ctx context.Context, email string) (GetEmailRow, error) {
	row := q.db.QueryRow(ctx, getEmail, email)
	var i GetEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getTokenPreRegister = `-- name: GetTokenPreRegister :one
SELECT register_token,
    token_expires_at
FROM users
WHERE register_token = $1
`

type GetTokenPreRegisterRow struct {
	RegisterToken  pgtype.Text
	TokenExpiresAt pgtype.Timestamp
}

func (q *Queries) GetTokenPreRegister(ctx context.Context, registerToken pgtype.Text) (GetTokenPreRegisterRow, error) {
	row := q.db.QueryRow(ctx, getTokenPreRegister, registerToken)
	var i GetTokenPreRegisterRow
	err := row.Scan(&i.RegisterToken, &i.TokenExpiresAt)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT users.id, users.name, users.email, users.document_type, users.document_number, users.password, users.status, users.register_token, users.token_expires_at, users.created_at, users.updated_at, users.tenant_id, users.image_id, address.id, address.user_id, address.address, address.street, address.city, address.state, address.postal_code, address.country, address.number, address.created_at, address.updated_at, image.id, image.description, image.url, image.created_at, image.updated_at, user_phones.id, user_phones.user_id, user_phones.type, user_phones.number, user_phones.is_primary, user_phones.created_at, user_phones.updated_at
FROM users
LEFT JOIN address ON address.user_id = users.id
LEFT JOIN image ON image.id = users.image_id
LEFT JOIN user_phones ON user_phones.user_id = users.id
WHERE users.id = $1
`

type GetUserRow struct {
	ID             pgtype.UUID
	Name           pgtype.Text
	Email          string
	DocumentType   pgtype.Text
	DocumentNumber pgtype.Text
	Password       pgtype.Text
	Status         UserStatus
	RegisterToken  pgtype.Text
	TokenExpiresAt pgtype.Timestamp
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
	TenantID       string
	ImageID        string
	Address        Address
	Image          Image
	UserPhone      UserPhone
}

func (q *Queries) GetUser(ctx context.Context, id pgtype.UUID) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.Password,
		&i.Status,
		&i.RegisterToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TenantID,
		&i.ImageID,
		&i.Address.ID,
		&i.Address.UserID,
		&i.Address.Address,
		&i.Address.Street,
		&i.Address.City,
		&i.Address.State,
		&i.Address.PostalCode,
		&i.Address.Country,
		&i.Address.Number,
		&i.Address.CreatedAt,
		&i.Address.UpdatedAt,
		&i.Image.ID,
		&i.Image.Description,
		&i.Image.Url,
		&i.Image.CreatedAt,
		&i.Image.UpdatedAt,
		&i.UserPhone.ID,
		&i.UserPhone.UserID,
		&i.UserPhone.Type,
		&i.UserPhone.Number,
		&i.UserPhone.IsPrimary,
		&i.UserPhone.CreatedAt,
		&i.UserPhone.UpdatedAt,
	)
	return i, err
}

const getUserContactEmail = `-- name: GetUserContactEmail :one
SELECT email, name
FROM contact_info
WHERE email = $1
`

type GetUserContactEmailRow struct {
	Email string
	Name  string
}

func (q *Queries) GetUserContactEmail(ctx context.Context, email string) (GetUserContactEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserContactEmail, email)
	var i GetUserContactEmailRow
	err := row.Scan(&i.Email, &i.Name)
	return i, err
}

const getUserRegisterToken = `-- name: GetUserRegisterToken :one
SELECT id, name, register_token
FROM users
WHERE register_token = $1
`

type GetUserRegisterTokenRow struct {
	ID            pgtype.UUID
	Name          pgtype.Text
	RegisterToken pgtype.Text
}

func (q *Queries) GetUserRegisterToken(ctx context.Context, registerToken pgtype.Text) (GetUserRegisterTokenRow, error) {
	row := q.db.QueryRow(ctx, getUserRegisterToken, registerToken)
	var i GetUserRegisterTokenRow
	err := row.Scan(&i.ID, &i.Name, &i.RegisterToken)
	return i, err
}

const getUsersContact = `-- name: GetUsersContact :many
SELECT email
FROM contact_info
`

func (q *Queries) GetUsersContact(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getUsersContact)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersWithEmail = `-- name: GetUsersWithEmail :many
SELECT email
FROM users
`

func (q *Queries) GetUsersWithEmail(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getUsersWithEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, document_type, document_number, password, status, register_token, token_expires_at, created_at, updated_at, tenant_id, image_id
FROM users
ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.DocumentType,
			&i.DocumentNumber,
			&i.Password,
			&i.Status,
			&i.RegisterToken,
			&i.TokenExpiresAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TenantID,
			&i.ImageID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $1,
    email = $2,
    document_type = $3,
    document_number = $4
WHERE id = $5
`

type UpdateUserParams struct {
	Name           pgtype.Text
	Email          string
	DocumentType   pgtype.Text
	DocumentNumber pgtype.Text
	ID             pgtype.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.ID,
	)
	return err
}
