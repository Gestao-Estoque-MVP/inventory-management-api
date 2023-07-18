// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const completeRegisterUser = `-- name: CompleteRegisterUser :one
UPDATE users SET phone = $1, document_type = $2, document_number = $3, password = $4, avatar = $5 WHERE id = $6 RETURNING id
`

type CompleteRegisterUserParams struct {
	Phone          sql.NullString
	DocumentType   sql.NullString
	DocumentNumber sql.NullString
	Password       sql.NullString
	Avatar         sql.NullString
	ID             string
}

func (q *Queries) CompleteRegisterUser(ctx context.Context, arg CompleteRegisterUserParams) (string, error) {
	row := q.db.QueryRowContext(ctx, completeRegisterUser,
		arg.Phone,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.Password,
		arg.Avatar,
		arg.ID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const createAddress = `-- name: CreateAddress :one
INSERT INTO address (user_id, address, number, street, city, state, postal_code, country) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, user_id, address, number, street, city, state, postal_code, country, created_at
`

type CreateAddressParams struct {
	UserID     string
	Address    sql.NullString
	Number     sql.NullString
	Street     sql.NullString
	City       sql.NullString
	State      sql.NullString
	PostalCode sql.NullString
	Country    sql.NullString
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, createAddress,
		arg.UserID,
		arg.Address,
		arg.Number,
		arg.Street,
		arg.City,
		arg.State,
		arg.PostalCode,
		arg.Country,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Address,
		&i.Number,
		&i.Street,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const createContactInfo = `-- name: CreateContactInfo :one
INSERT INTO contact_info (id, name, email, phone, created_at) 
    VALUES ($1, $2, $3, $4, $5) RETURNING id, name, email, phone, created_at
`

type CreateContactInfoParams struct {
	ID        string
	Name      string
	Email     string
	Phone     sql.NullString
	CreatedAt time.Time
}

func (q *Queries) CreateContactInfo(ctx context.Context, arg CreateContactInfoParams) (ContactInfo, error) {
	row := q.db.QueryRowContext(ctx, createContactInfo,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.CreatedAt,
	)
	var i ContactInfo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.CreatedAt,
	)
	return i, err
}

const createPermissions = `-- name: CreatePermissions :one
INSERT INTO permissions (id, name, description) 
    VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreatePermissionsParams struct {
	ID          string
	Name        string
	Description string
}

func (q *Queries) CreatePermissions(ctx context.Context, arg CreatePermissionsParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, createPermissions, arg.ID, arg.Name, arg.Description)
	var i Permission
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createPreRegisterUser = `-- name: CreatePreRegisterUser :one
INSERT INTO users (id, name, email, status, role_id, register_token, token_expires_at, created_at) 
    VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, name, email
`

type CreatePreRegisterUserParams struct {
	ID             string
	Name           string
	Email          string
	Status         string
	RoleID         sql.NullString
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
		arg.RegisterToken,
		arg.TokenExpiresAt,
		arg.CreatedAt,
	)
	var i CreatePreRegisterUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles (id, name, description) 
    VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateRoleParams struct {
	ID          string
	Name        string
	Description string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, arg.ID, arg.Name, arg.Description)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createRolePermissions = `-- name: CreateRolePermissions :one
INSERT INTO roles_permissions (role_id, permission_id) 
    VALUES ($1, $2) RETURNING role_id, permission_id
`

type CreateRolePermissionsParams struct {
	RoleID       string
	PermissionID string
}

func (q *Queries) CreateRolePermissions(ctx context.Context, arg CreateRolePermissionsParams) (RolesPermission, error) {
	row := q.db.QueryRowContext(ctx, createRolePermissions, arg.RoleID, arg.PermissionID)
	var i RolesPermission
	err := row.Scan(&i.RoleID, &i.PermissionID)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :execresult
DELETE FROM address WHERE id = $1 RETURNING id, user_id, address, number, street, city, state, postal_code, country, created_at
`

func (q *Queries) DeleteAddress(ctx context.Context, id int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAddress, id)
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

const getAddress = `-- name: GetAddress :one
SELECT id, user_id, address, number, street, city, state, postal_code, country, created_at FROM address WHERE id = $1
`

func (q *Queries) GetAddress(ctx context.Context, id int32) (Address, error) {
	row := q.db.QueryRowContext(ctx, getAddress, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Address,
		&i.Number,
		&i.Street,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const getRole = `-- name: GetRole :one
SELECT 
    u.id AS user_id,
    r.name 
FROM 
    users AS u
INNER JOIN 
    roles AS r ON u.role_id = r.id
WHERE 
    u.id = $1
`

type GetRoleRow struct {
	UserID string
	Name   string
}

func (q *Queries) GetRole(ctx context.Context, id string) (GetRoleRow, error) {
	row := q.db.QueryRowContext(ctx, getRole, id)
	var i GetRoleRow
	err := row.Scan(&i.UserID, &i.Name)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, phone, document_type, document_number, password, avatar, status, register_token, token_expires_at, created_at, role_id FROM users WHERE id = $1
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
		&i.RoleID,
	)
	return i, err
}

const listAddresses = `-- name: ListAddresses :many
SELECT id, user_id, address, number, street, city, state, postal_code, country, created_at FROM address
`

func (q *Queries) ListAddresses(ctx context.Context) ([]Address, error) {
	rows, err := q.db.QueryContext(ctx, listAddresses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Address
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Address,
			&i.Number,
			&i.Street,
			&i.City,
			&i.State,
			&i.PostalCode,
			&i.Country,
			&i.CreatedAt,
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

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, phone, document_type, document_number, password, avatar, status, register_token, token_expires_at, created_at, role_id FROM users
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
			&i.RoleID,
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

const updateAddress = `-- name: UpdateAddress :one
UPDATE address SET user_id = $1, address = $2, number = $3, street = $4, city = $5, state = $6, postal_code = $7, country = $8 WHERE id = $9 RETURNING id, user_id, address, number, street, city, state, postal_code, country, created_at
`

type UpdateAddressParams struct {
	UserID     string
	Address    sql.NullString
	Number     sql.NullString
	Street     sql.NullString
	City       sql.NullString
	State      sql.NullString
	PostalCode sql.NullString
	Country    sql.NullString
	ID         int32
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, updateAddress,
		arg.UserID,
		arg.Address,
		arg.Number,
		arg.Street,
		arg.City,
		arg.State,
		arg.PostalCode,
		arg.Country,
		arg.ID,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Address,
		&i.Number,
		&i.Street,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET name = $1, email = $2, phone = $3, document_type = $4, document_number = $5, password = $6, status = $7 WHERE id = $8 RETURNING id, name, email
`

type UpdateUserParams struct {
	Name           string
	Email          string
	Phone          sql.NullString
	DocumentType   sql.NullString
	DocumentNumber sql.NullString
	Password       sql.NullString
	Status         string
	ID             string
}

type UpdateUserRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.Password,
		arg.Status,
		arg.ID,
	)
	return err
}
