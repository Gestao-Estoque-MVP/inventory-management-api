package service

import (
	"math/big"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/graph/model"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/diogoX451/inventory-management-api/pkg/convert"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateCategory(category database.CreateCategoryParams) (*pgtype.UUID, error) {
	id, _ := uuid.NewV4()
	params := database.CreateCategoryParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
	}

	create, err := s.repo.CreateCategory(params)

	if err != nil {
		return nil, err
	}

	return create, err
}

func (s *ProductService) CreateProductUnitsOfMeasure(units database.CreateProductUnitsOfMeasureParams) (*pgtype.UUID, error) {
	id, _ := uuid.NewV4()

	params := database.CreateProductUnitsOfMeasureParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		Name:        units.Name,
		Description: units.Description,
		CreatedAt:   pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
	}

	create, err := s.repo.CreateProductUnitsOfMeasure(params)

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (s *ProductService) CreateVariationsCategories(variations database.CreateVariationsCategoriesParams) (*pgtype.UUID, error) {
	id, _ := uuid.NewV4()

	params := database.CreateVariationsCategoriesParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		Name:        variations.Name,
		Description: variations.Description,
		CreatedAt:   pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
	}

	create, err := s.repo.CreateVariationsCategories(params)

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (s *ProductService) CreateProduct(product model.NewProductInput) (*pgtype.UUID, error) {
	id, _ := uuid.NewV4()

	params := database.CreateProductParams{
		ID:                     pgtype.UUID{Bytes: id, Valid: true},
		Name:                   product.Name,
		LowStockThreshold:      pgtype.Int4{Int32: int32(*product.LowStockThreshold), Valid: true},
		Price:                  product.Price,
		Promotion:              pgtype.Float8{Float64: *product.Promotion, Valid: true},
		TenantID:               convert.StringToPgUUID(*product.TenantID),
		SafetyStockLevel:       pgtype.Int4{Int32: int32(*product.SafetyStockLevel), Valid: true},
		ReorderPoint:           pgtype.Int4{Int32: int32(*product.ReorderPoint), Valid: true},
		MinLot:                 pgtype.Int4{Int32: int32(*product.MinLot), Valid: true},
		MaxLot:                 pgtype.Int4{Int32: int32(*product.MaxLot), Valid: true},
		Width:                  pgtype.Int4{Int32: int32(*product.Width), Valid: true},
		Height:                 pgtype.Int4{Int32: int32(*product.Height), Valid: true},
		Length:                 pgtype.Int4{Int32: int32(*product.Length), Valid: true},
		Weight:                 pgtype.Int4{Int32: int32(*product.Weight), Valid: true},
		ProductUnitOfMeasureID: convert.StringToPgUUID(*product.ProductUnitOfMeasureID),
		IsVariation:            pgtype.Bool{Bool: *product.IsVariation, Valid: true},
		FsnClassification:      pgtype.Text{String: *product.FsnClassification, Valid: true},
		IsActive:               pgtype.Bool{Bool: true, Valid: true},
		CreatedAt:              pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
	}

	create, err := s.repo.CreateProduct(params)

	if err != nil {
		return nil, err
	}

	errors := make(chan error)

	go func(productID pgtype.UUID, items []*model.NewItemsVariations) {
		var values []model.NewItemsVariations
		if *product.IsVariation {
			for _, item := range items {
				values = append(values, *item)
			}

			err := variations(values, s.repo, productID)
			if err != nil {
				errors <- err
			}
		}
	}(*create, product.ItemsVariations)

	return create, nil
}

func variations(items []model.NewItemsVariations, repo repository.ProductRepository, productID pgtype.UUID) error {
	groupedVariations := make(map[string][]model.NewItemsVariations)
	for _, v := range items {
		groupedVariations[v.VaritationCategoryID] = append(groupedVariations[v.VaritationCategoryID], v)
	}

	possibleVariations := [][]model.NewItemsVariations{{}}

	for _, group := range groupedVariations {
		newCombinations := [][]model.NewItemsVariations{}
		for _, existingCombination := range possibleVariations {
			for _, variation := range group {
				newCombination := append([]model.NewItemsVariations(nil), existingCombination...)
				newCombination = append(newCombination, variation)
				newCombinations = append(newCombinations, newCombination)
			}
		}
		possibleVariations = newCombinations
	}

	createVariations(possibleVariations, repo, productID)

	return nil
}

func createVariations(variations [][]model.NewItemsVariations, repo repository.ProductRepository, productID pgtype.UUID) error {

	variationsIds := make(map[string]pgtype.UUID)
	for i := 0; i < len(variations); i++ {
		params := database.CreateProductVariationsParams{
			ID:        pgtype.UUID{Bytes: uuid.Must(uuid.NewV4()), Valid: true},
			ProductID: productID,
			Price:     pgtype.Numeric{Int: big.NewInt(0), Valid: true},
			Promotion: pgtype.Numeric{Int: big.NewInt(0), Valid: true},
			CreatedAt: pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
		}

		create, err := repo.CreateProductVariatons(params)

		if err != nil {
			return err
		}

		variationsIds
	}

	return nil

}
