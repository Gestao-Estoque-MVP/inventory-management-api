// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: address_queries.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO address (user_id, address, number, street, city, state, postal_code, country, created_at) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at
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
	CreatedAt  time.Time
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
		arg.CreatedAt,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Address,
		&i.Street,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.Number,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :execresult
DELETE FROM address WHERE user_id = $1 RETURNING id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at
`

func (q *Queries) DeleteAddress(ctx context.Context, userID string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAddress, userID)
}

const getAddressByID = `-- name: GetAddressByID :one
SELECT id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at FROM address WHERE user_id = $1
`

func (q *Queries) GetAddressByID(ctx context.Context, userID string) (Address, error) {
	row := q.db.QueryRowContext(ctx, getAddressByID, userID)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Address,
		&i.Street,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.Number,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAddresses = `-- name: ListAddresses :many
SELECT id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at FROM address
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
			&i.Street,
			&i.City,
			&i.State,
			&i.PostalCode,
			&i.Country,
			&i.Number,
			&i.CreatedAt,
			&i.UpdatedAt,
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
UPDATE address SET user_id = $1, address = $2, number = $3, street = $4, city = $5, state = $6, postal_code = $7, country = $8 WHERE id = $9 RETURNING id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at
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
		&i.Street,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.Number,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
