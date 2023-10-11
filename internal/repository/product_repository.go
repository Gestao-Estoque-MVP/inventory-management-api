package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type IProduct interface {
	CreateProduct(product database.CreateProductParams) (*pgtype.UUID, error)
	CreateCategory(category database.CreateCategoryParams) (*pgtype.UUID, error)
	CreateProductCategory(prodCat database.CreateProductsCategoriesParams) (*pgtype.UUID, error)
}

type ProductRepository struct {
	db *database.Queries
}

func NewProductRepository(db *database.Queries) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) CreateProduct(product database.CreateProductParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateProduct(context.Background(), product)

	if err != nil {
		return nil, err
	}

	return &create, err
}

func (repo *ProductRepository) CreateCategory(category database.CreateCategoryParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateCategory(context.Background(), category)

	if err == nil {
		return nil, err
	}

	return &create, nil
}

func (repo *ProductRepository) CreateProductsCategories(prodCat database.CreateProductsCategoriesParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateProductsCategories(context.Background(), prodCat)
	if err == nil {
		return nil, err
	}

	return &create, nil
}
