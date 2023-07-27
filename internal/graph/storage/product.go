package storage

import (
	"context"
	"database/sql"

	"github.com/graph-gophers/dataloader"
)

type contextKey string

const (
	loaders = contextKey("products-loaders")
)

type ProductReader struct {
	db *sql.DB
}

func (u *ProductReader) GetProduct(ctx context.Context, keys dataloader.Keys) ([]*dataloader.Result, error) {
}
