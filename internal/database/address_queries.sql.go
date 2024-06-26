// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: address_queries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO address (id, user_id, address, number, street, city, state, postal_code, country, created_at) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at
`

type CreateAddressParams struct {
	ID         pgtype.UUID
	UserID     pgtype.UUID
	Address    pgtype.Text
	Number     pgtype.Text
	Street     pgtype.Text
	City       pgtype.Text
	State      pgtype.Text
	PostalCode pgtype.Text
	Country    pgtype.Text
	CreatedAt  pgtype.Timestamp
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error) {
	row := q.db.QueryRow(ctx, createAddress,
		arg.ID,
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
DELETE FROM address WHERE user_id = $1
`

func (q *Queries) DeleteAddress(ctx context.Context, userID pgtype.UUID) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteAddress, userID)
}

const getAddressByID = `-- name: GetAddressByID :one
SELECT id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at FROM address WHERE user_id = $1
`

func (q *Queries) GetAddressByID(ctx context.Context, userID pgtype.UUID) (Address, error) {
	row := q.db.QueryRow(ctx, getAddressByID, userID)
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
	rows, err := q.db.Query(ctx, listAddresses)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAddress = `-- name: UpdateAddress :one
UPDATE address SET user_id = $1, address = $2, number = $3, street = $4, city = $5, state = $6, postal_code = $7, country = $8 WHERE id = $9 RETURNING id, user_id, address, street, city, state, postal_code, country, number, created_at, updated_at
`

type UpdateAddressParams struct {
	UserID     pgtype.UUID
	Address    pgtype.Text
	Number     pgtype.Text
	Street     pgtype.Text
	City       pgtype.Text
	State      pgtype.Text
	PostalCode pgtype.Text
	Country    pgtype.Text
	ID         pgtype.UUID
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error) {
	row := q.db.QueryRow(ctx, updateAddress,
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
