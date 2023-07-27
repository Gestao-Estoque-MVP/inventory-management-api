package storage

import (
	"context"
	"database/sql"

	"github.com/diogoX451/inventory-management-api/internal/service"
	"github.com/graph-gophers/dataloader"
)

type contextKey string

const (
	loaders = contextKey("users-loaders")
)

type UserReader struct {
	db *sql.DB
}

func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.Keys) ([]*dataloader.Result, error) {
	W
}

func UserBatchFn(userService *service.UserService) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := keys.Keys()
		for _, id := range ids {
			user, err := userService.GetUser(id)
			if err != nil {
				results = append(results, &dataloader.Result{Error: err})
				continue
			}

			results = append(results, &dataloader.Result{Data: user})
		}

		return results
	}
}

// AddressBatchFn busca endereços por seus IDs
func AddressBatchFn(addressService *service.AddressService) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := keys.Keys()
		for _, id := range ids {
			address, err := addressService.GetAddressByID(id)
			if err != nil {
				results = append(results, &dataloader.Result{Error: err})
				continue
			}

			results = append(results, &dataloader.Result{Data: address})
		}

		return results
	}
}
