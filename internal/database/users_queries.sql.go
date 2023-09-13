// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: users_queries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const completeRegisterUser = `-- name: CompleteRegisterUser :one
UPDATE users
SET document_type = $1,
    document_number = $2,
    password = $3,
    status = $4,
    image_id = $5,
    updated_at = $6
WHERE register_token = $7
RETURNING id,
    name,
    email
`

type CompleteRegisterUserParams struct {
	DocumentType   pgtype.Text
	DocumentNumber pgtype.Text
	Password       pgtype.Text
	Status         UserStatus
	ImageID        string
	UpdatedAt      pgtype.Timestamp
	RegisterToken  pgtype.Text
}

type CompleteRegisterUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) CompleteRegisterUser(ctx context.Context, arg CompleteRegisterUserParams) (CompleteRegisterUserRow, error) {
	row := q.db.QueryRow(ctx, completeRegisterUser,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.Password,
		arg.Status,
		arg.ImageID,
		arg.UpdatedAt,
		arg.RegisterToken,
	)
	var i CompleteRegisterUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createPreRegisterUser = `-- name: CreatePreRegisterUser :one
INSERT INTO users (
    id, 
    name, 
    email, 
    status, 
    register_token, 
    token_expires_at, 
    created_at, 
    role_id, 
    tenant_id
) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, email
`

type CreatePreRegisterUserParams struct {
	ID             string
	Name           string
	Email          string
	Status         UserStatus
	RegisterToken  pgtype.Text
	TokenExpiresAt pgtype.Timestamp
	CreatedAt      pgtype.Timestamp
	RoleID         pgtype.Text
	TenantID       string
}

type CreatePreRegisterUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) CreatePreRegisterUser(ctx context.Context, arg CreatePreRegisterUserParams) (CreatePreRegisterUserRow, error) {
	row := q.db.QueryRow(ctx, createPreRegisterUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Status,
		arg.RegisterToken,
		arg.TokenExpiresAt,
		arg.CreatedAt,
		arg.RoleID,
		arg.TenantID,
	)
	var i CreatePreRegisterUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenant (id, name, tax_code, type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, tax_code, type, created_at, updated_at
`

type CreateTenantParams struct {
	ID        string
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
    created_at,
    updated_at
) VALUES($1, $2, $3, $4,$5,$6) RETURNING id, number, type
`

type CreateUserPhonesParams struct {
	ID        string
	Type      TypeNumber
	Number    string
	IsPrimary bool
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type CreateUserPhonesRow struct {
	ID     string
	Number string
	Type   TypeNumber
}

func (q *Queries) CreateUserPhones(ctx context.Context, arg CreateUserPhonesParams) (CreateUserPhonesRow, error) {
	row := q.db.QueryRow(ctx, createUserPhones,
		arg.ID,
		arg.Type,
		arg.Number,
		arg.IsPrimary,
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
	ID    string
	Name  string
	Email string
}

func (q *Queries) DeleteUser(ctx context.Context, id string) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteUser, id)
}

const getEmail = `-- name: GetEmail :one
SELECT id,
    name,
    email,
    password,
    role_id
FROM users
WHERE email = $1
`

type GetEmailRow struct {
	ID       string
	Name     string
	Email    string
	Password pgtype.Text
	RoleID   pgtype.Text
}

func (q *Queries) GetEmail(ctx context.Context, email string) (GetEmailRow, error) {
	row := q.db.QueryRow(ctx, getEmail, email)
	var i GetEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.RoleID,
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
SELECT id, name, email, document_type, document_number, password, status, register_token, token_expires_at, created_at, updated_at, role_id, tenant_id, image_id
FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
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
		&i.RoleID,
		&i.TenantID,
		&i.ImageID,
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
SELECT id, name, email, document_type, document_number, password, status, register_token, token_expires_at, created_at, updated_at, role_id, tenant_id, image_id
FROM users
WHERE register_token = $1
`

func (q *Queries) GetUserRegisterToken(ctx context.Context, registerToken pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserRegisterToken, registerToken)
	var i User
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
		&i.RoleID,
		&i.TenantID,
		&i.ImageID,
	)
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
SELECT id, name, email, document_type, document_number, password, status, register_token, token_expires_at, created_at, updated_at, role_id, tenant_id, image_id
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
			&i.RoleID,
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
	Name           string
	Email          string
	DocumentType   pgtype.Text
	DocumentNumber pgtype.Text
	ID             string
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
