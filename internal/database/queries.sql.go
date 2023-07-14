// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO address (user_id, address, number, street, city, state, postal_code, country) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, user_id, address, number, street, city, state, postal_code, country
`

type CreateAddressParams struct {
	UserID     sql.NullInt32
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
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, lastname, email, phone, document_type, document_number, password ) 
    VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, lastname, email, phone, document_type, document_number, password
`

type CreateUserParams struct {
	Name           string
	Lastname       string
	Email          string
	Phone          string
	DocumentType   string
	DocumentNumber string
	Password       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Lastname,
		arg.Email,
		arg.Phone,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Phone,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.Password,
	)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :one
DELETE FROM address WHERE id = $1 RETURNING id, user_id, address, number, street, city, state, postal_code, country
`

func (q *Queries) DeleteAddress(ctx context.Context, id int32) (Address, error) {
	row := q.db.QueryRowContext(ctx, deleteAddress, id)
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
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING id, name, lastname, email, phone, document_type, document_number, password
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Phone,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.Password,
	)
	return i, err
}

const getAddress = `-- name: GetAddress :one
SELECT id, user_id, address, number, street, city, state, postal_code, country FROM address WHERE id = $1
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
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, lastname, email, phone, document_type, document_number, password FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Phone,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.Password,
	)
	return i, err
}

const listAddresses = `-- name: ListAddresses :many
SELECT id, user_id, address, number, street, city, state, postal_code, country FROM address
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
SELECT id, name, lastname, email, phone, document_type, document_number, password FROM users
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
			&i.Lastname,
			&i.Email,
			&i.Phone,
			&i.DocumentType,
			&i.DocumentNumber,
			&i.Password,
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
