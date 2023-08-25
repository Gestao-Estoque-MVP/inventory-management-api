// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users_queries.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const completeRegisterUser = `-- name: CompleteRegisterUser :one
UPDATE users SET phone = $1, document_type = $2, document_number = $3, password = $4, avatar = $5, updated_at = $6 WHERE register_token = $7 RETURNING id, name, email
`

type CompleteRegisterUserParams struct {
	Phone          sql.NullString
	DocumentType   sql.NullString
	DocumentNumber sql.NullString
	Password       sql.NullString
	Avatar         sql.NullString
	UpdatedAt      sql.NullTime
	RegisterToken  sql.NullString
}

type CompleteRegisterUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) CompleteRegisterUser(ctx context.Context, arg CompleteRegisterUserParams) (CompleteRegisterUserRow, error) {
	row := q.db.QueryRowContext(ctx, completeRegisterUser,
		arg.Phone,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.Password,
		arg.Avatar,
		arg.UpdatedAt,
		arg.RegisterToken,
	)
	var i CompleteRegisterUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createPreRegisterUser = `-- name: CreatePreRegisterUser :one
INSERT INTO users (id, name, email, status, role_id, tenant_id, register_token, token_expires_at, created_at) 
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, email
`

type CreatePreRegisterUserParams struct {
	ID             string
	Name           string
	Email          string
	Status         UserStatus
	RoleID         sql.NullString
	TenantID       string
	RegisterToken  sql.NullString
	TokenExpiresAt sql.NullTime
	CreatedAt      time.Time
}

type CreatePreRegisterUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) CreatePreRegisterUser(ctx context.Context, arg CreatePreRegisterUserParams) (CreatePreRegisterUserRow, error) {
	row := q.db.QueryRowContext(ctx, createPreRegisterUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Status,
		arg.RoleID,
		arg.TenantID,
		arg.RegisterToken,
		arg.TokenExpiresAt,
		arg.CreatedAt,
	)
	var i CreatePreRegisterUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenant (id, name) 
    VALUES ($1, $2) RETURNING id, name
`

type CreateTenantParams struct {
	ID   string
	Name sql.NullString
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, createTenant, arg.ID, arg.Name)
	var i Tenant
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteUser = `-- name: DeleteUser :execresult
DELETE FROM users WHERE id = $1 RETURNING id, name, email
`

type DeleteUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) DeleteUser(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, id)
}

const getEmail = `-- name: GetEmail :one
SELECT id, name, email, password, role_id FROM users WHERE email = $1
`

type GetEmailRow struct {
	ID       string
	Name     string
	Email    string
	Password sql.NullString
	RoleID   sql.NullString
}

func (q *Queries) GetEmail(ctx context.Context, email string) (GetEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getEmail, email)
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
SELECT register_token FROM users WHERE register_token = $1
`

func (q *Queries) GetTokenPreRegister(ctx context.Context, registerToken sql.NullString) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getTokenPreRegister, registerToken)
	var register_token sql.NullString
	err := row.Scan(&register_token)
	return register_token, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, phone, document_type, document_number, password, avatar, status, register_token, token_expires_at, created_at, updated_at, role_id, tenant_id FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.Password,
		&i.Avatar,
		&i.Status,
		&i.RegisterToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoleID,
		&i.TenantID,
	)
	return i, err
}

const getUserRegisterToken = `-- name: GetUserRegisterToken :one
SELECT id, name, email, phone, document_type, document_number, password, avatar, status, register_token, token_expires_at, created_at, updated_at, role_id, tenant_id FROM users WHERE register_token = $1
`

func (q *Queries) GetUserRegisterToken(ctx context.Context, registerToken sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserRegisterToken, registerToken)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.Password,
		&i.Avatar,
		&i.Status,
		&i.RegisterToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoleID,
		&i.TenantID,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, phone, document_type, document_number, password, avatar, status, register_token, token_expires_at, created_at, updated_at, role_id, tenant_id FROM users ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
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
			&i.Phone,
			&i.DocumentType,
			&i.DocumentNumber,
			&i.Password,
			&i.Avatar,
			&i.Status,
			&i.RegisterToken,
			&i.TokenExpiresAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RoleID,
			&i.TenantID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET name = $1, email = $2, phone = $3, document_type = $4, document_number = $5 WHERE id = $6
`

type UpdateUserParams struct {
	Name           string
	Email          string
	Phone          sql.NullString
	DocumentType   sql.NullString
	DocumentNumber sql.NullString
	ID             string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.ID,
	)
	return err
}
