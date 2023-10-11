package service

import (
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductService struct {
	repo repository.IProduct
}

func NewProductService(repo repository.IProduct) *ProductService {
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

func (s *ProductService) CreateProduct(product database.CreateProductParams) (*pgtype.UUID, error) {
	id, _ := uuid.NewV4()
	params := database.CreateProductParams{
		ID:                pgtype.UUID{Bytes: id, Valid: true},
		Name:              product.Name,
		LowStockThreshold: product.LowStockThreshold,
		Price:             product.Price,
		TenantID:          product.TenantID,
		Promotion:         product.Promotion,
		SafetyStockLevel:  product.SafetyStockLevel,
		ReorderPoint:      product.ReorderPoint,
		MinLot:            product.MinLot,
		MaxLot:            product.MaxLot,
		FsnClassification: product.FsnClassification,
		Width:             product.Width,
		Height:            product.Height,
		Length:            product.Length,
		Weight:            product.Weight,
	}

	create, err := s.repo.CreateProduct(params)

	if err != nil {
		return nil, err
	}

	go func(productID) {
		if  {
			
		}
	}(create)

	return create, nil
}
