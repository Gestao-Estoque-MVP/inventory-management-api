// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: auth_queries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createContactInfo = `-- name: CreateContactInfo :one
INSERT INTO contact_info (id, name, email, phone, created_at) 
    VALUES ($1, $2, $3, $4, $5) RETURNING id, name, email, phone, created_at
`

type CreateContactInfoParams struct {
	ID        pgtype.UUID
	Name      string
	Email     string
	Phone     pgtype.Text
	CreatedAt pgtype.Timestamp
}

func (q *Queries) CreateContactInfo(ctx context.Context, arg CreateContactInfoParams) (ContactInfo, error) {
	row := q.db.QueryRow(ctx, createContactInfo,
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
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreatePermissions(ctx context.Context, arg CreatePermissionsParams) (Permission, error) {
	row := q.db.QueryRow(ctx, createPermissions, arg.ID, arg.Name, arg.Description)
	var i Permission
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles (id, name, description) 
    VALUES ($1, $2, $3) RETURNING id, name, description
`

type CreateRoleParams struct {
	ID          pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, arg.ID, arg.Name, arg.Description)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createRolePermissions = `-- name: CreateRolePermissions :one
INSERT INTO roles_permissions (id, role_id, permission_id) 
    VALUES ($1, $2, $3) RETURNING id, role_id, permission_id
`

type CreateRolePermissionsParams struct {
	ID           pgtype.UUID
	RoleID       pgtype.UUID
	PermissionID pgtype.UUID
}

func (q *Queries) CreateRolePermissions(ctx context.Context, arg CreateRolePermissionsParams) (RolesPermission, error) {
	row := q.db.QueryRow(ctx, createRolePermissions, arg.ID, arg.RoleID, arg.PermissionID)
	var i RolesPermission
	err := row.Scan(&i.ID, &i.RoleID, &i.PermissionID)
	return i, err
}

const createUsersPermissions = `-- name: CreateUsersPermissions :one
INSERT INTO users_permissions (user_id, permission_id) 
    VALUES ($1, $2) RETURNING id, user_id, permission_id
`

type CreateUsersPermissionsParams struct {
	UserID       pgtype.UUID
	PermissionID pgtype.UUID
}

func (q *Queries) CreateUsersPermissions(ctx context.Context, arg CreateUsersPermissionsParams) (UsersPermission, error) {
	row := q.db.QueryRow(ctx, createUsersPermissions, arg.UserID, arg.PermissionID)
	var i UsersPermission
	err := row.Scan(&i.ID, &i.UserID, &i.PermissionID)
	return i, err
}

const createUsersRoles = `-- name: CreateUsersRoles :one

INSERT INTO users_roles (id, user_id, role_id) 
    VALUES ($1, $2, $3) RETURNING id, user_id, role_id
`

type CreateUsersRolesParams struct {
	ID     pgtype.UUID
	UserID pgtype.UUID
	RoleID pgtype.UUID
}

func (q *Queries) CreateUsersRoles(ctx context.Context, arg CreateUsersRolesParams) (UsersRole, error) {
	row := q.db.QueryRow(ctx, createUsersRoles, arg.ID, arg.UserID, arg.RoleID)
	var i UsersRole
	err := row.Scan(&i.ID, &i.UserID, &i.RoleID)
	return i, err
}

const getRole = `-- name: GetRole :one
SELECT id, name, description FROM roles WHERE name = $1
`

func (q *Queries) GetRole(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getRoleByID = `-- name: GetRoleByID :one
SELECT id, name, description from roles WHERE id = $1
`

func (q *Queries) GetRoleByID(ctx context.Context, id pgtype.UUID) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByID, id)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getRoleUser = `-- name: GetRoleUser :one
SELECT r.id, r.name
FROM users u
JOIN users_roles ur ON u.id = ur.user_id
JOIN roles r ON ur.role_id = r.id
WHERE u.id = $1
`

type GetRoleUserRow struct {
	ID   pgtype.UUID
	Name string
}

func (q *Queries) GetRoleUser(ctx context.Context, id pgtype.UUID) (GetRoleUserRow, error) {
	row := q.db.QueryRow(ctx, getRoleUser, id)
	var i GetRoleUserRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getRolesPermissions = `-- name: GetRolesPermissions :many
SELECT
    u.id AS user_id,
    r.name AS role_name,
    p.name AS permission_name
FROM 
    users AS u
INNER JOIN 
    roles AS r ON u.role_id = r.id
INNER JOIN 
    roles_permissions AS rp ON r.id = rp.role_id
INNER JOIN
    permissions AS p ON rp.permission_id = p.id
WHERE 
    u.id = $1
`

type GetRolesPermissionsRow struct {
	UserID         pgtype.UUID
	RoleName       string
	PermissionName string
}

func (q *Queries) GetRolesPermissions(ctx context.Context, id pgtype.UUID) ([]GetRolesPermissionsRow, error) {
	rows, err := q.db.Query(ctx, getRolesPermissions, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRolesPermissionsRow
	for rows.Next() {
		var i GetRolesPermissionsRow
		if err := rows.Scan(&i.UserID, &i.RoleName, &i.PermissionName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersPermissions = `-- name: GetUsersPermissions :many
SELECT
    u.id AS user_id,
    p.name AS permission_name
FROM
    users AS u
INNER JOIN
    users_permissions AS up ON u.id = up.user_id
INNER JOIN
    permissions AS p ON up.permission_id = p.id
WHERE
    u.id = $1
`

type GetUsersPermissionsRow struct {
	UserID         pgtype.UUID
	PermissionName string
}

func (q *Queries) GetUsersPermissions(ctx context.Context, id pgtype.UUID) ([]GetUsersPermissionsRow, error) {
	rows, err := q.db.Query(ctx, getUsersPermissions, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersPermissionsRow
	for rows.Next() {
		var i GetUsersPermissionsRow
		if err := rows.Scan(&i.UserID, &i.PermissionName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
