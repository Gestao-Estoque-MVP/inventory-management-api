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
	CreateProductUnitsOfMeasur(product database.CreateProductUnitsOfMeasureParams) (*pgtype.UUID, error)
	CreateVariationsCategories(variations database.CreateVariationsCategoriesParams) (*pgtype.UUID, error)
	CreateVariationsItems(variations database.CreateVariationsItemsParams) (*pgtype.UUID, error)
	CreateProductVariatons(product database.CreateProductVariationsParams) (*pgtype.UUID, error)
	CreateVariationsMappings(variations database.CreateVariationMappingParams) (*pgtype.UUID, error)
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

func (repo *ProductRepository) CreateProductUnitsOfMeasure(product database.CreateProductUnitsOfMeasureParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateProductUnitsOfMeasure(context.Background(), product)

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (repo *ProductRepository) CreateVariationsCategories(variations database.CreateVariationsCategoriesParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateVariationsCategories(context.Background(), variations)

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (repo *ProductRepository) CreateVariationsItems(variations database.CreateVariationsItemsParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateVariationsItems(context.Background(), variations)

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (repo *ProductRepository) CreateProductVariatons(product database.CreateProductVariationsParams) (*pgtype.UUID, error) {
	create, err := repo.db.CreateProductVariations(context.Background(), product)

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (repo *ProductRepository) CreateVariationsMappings(variations database.CreateVariationMappingParams) *pgtype.UUID {
	create, err := repo.db.CreateVariationMapping(context.Background(), variations)

	if err != nil {
		return nil
	}

	return &create
}
