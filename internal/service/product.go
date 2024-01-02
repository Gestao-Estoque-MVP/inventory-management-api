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
	var mapItems []model.ItemProduct

	for i := 0; i < len(items); i++ {
		create, _ := repo.CreateVariationsItems(database.CreateVariationsItemsParams{
			ID:                  pgtype.UUID{Bytes: uuid.Must(uuid.NewV4()), Valid: true},
			VariationCategoryID: convert.StringToPgUUID(items[i].VaritationCategoryID),
			Name:                items[i].Name,
			CreatedAt:           pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
		})

		mapItems = append(mapItems, model.ItemProduct{
			ID:   convert.UUIDToString(*create),
			Name: items[i].Name,
		})
	}

	possible := foreachVariations(items, productID)

	for _, combo := range possible {
		create, _ := repo.CreateProductVariatons(database.CreateProductVariationsParams{
			ID:        pgtype.UUID{Bytes: uuid.Must(uuid.NewV4()), Valid: true},
			ProductID: productID,
			Price:     pgtype.Numeric{Int: big.NewInt(0), Valid: true},
			Promotion: pgtype.Numeric{Int: big.NewInt(0), Valid: true},
			CreatedAt: pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
		})
		variationID := *create

		for _, item := range combo {
			var itemID string
			for _, mapItem := range mapItems {
				if item.Name == mapItem.Name {
					itemID = mapItem.ID
					break
				}
			}

			repo.CreateVariationsMappings(database.CreateVariationMappingParams{
				ID:                 pgtype.UUID{Bytes: uuid.Must(uuid.NewV4()), Valid: true},
				VariationItemID:    convert.StringToPgUUID(itemID),
				ProductVariationID: variationID,
				CreatedAt:          pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
			})
		}
	}

	return nil
}

func foreachVariations(items []model.NewItemsVariations, productID pgtype.UUID) [][]model.NewItemsVariations {
	categoryMap := make(map[string][]model.NewItemsVariations)

	for _, v := range items {
		categoryMap[v.VaritationCategoryID] = append(categoryMap[v.VaritationCategoryID], v)
	}

	var result [][]model.NewItemsVariations
	result = append(result, []model.NewItemsVariations{})

	for _, cm := range categoryMap {
		var temp [][]model.NewItemsVariations
		for _, combo := range result {
			for _, v := range cm {
				newCombo := make([]model.NewItemsVariations, len(combo))
				copy(newCombo, combo)
				newCombo = append(newCombo, v)
				temp = append(temp, newCombo)
			}
		}
		result = temp
	}

	return result

}

// func createVariations(variations [][]model.NewItemsVariations, repo repository.ProductRepository, productID pgtype.UUID) error {

// 	for i := 0; i < len(variations); i++ {
// 		params := database.CreateProductVariationsParams{
// 			ID:        pgtype.UUID{Bytes: uuid.Must(uuid.NewV4()), Valid: true},
// 			ProductID: productID,
// 			Price:     pgtype.Numeric{Int: big.NewInt(0), Valid: true},
// 			Promotion: pgtype.Numeric{Int: big.NewInt(0), Valid: true},
// 			CreatedAt: pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
// 		}

// 		_, err := repo.CreateProductVariatons(params)

// 		if err != nil {
// 			return err
// 		}

// 	}

// 	return nil

// }
