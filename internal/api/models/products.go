package models

import (
	"time"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	"gorm.io/gorm"
)

type Product struct {
	ID           uint64         `gorm:"column:id;primaryKey"`
	ProductCode  string         `gorm:"column:product_code"`
	ProductName  string         `gorm:"column:product_name"`
	ProductPrice float64        `gorm:"column:product_price"`
	CreatedBy    string         `gorm:"column:created_by"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedBy    string         `gorm:"column:updated_by"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (m Product) Entity() *entities.Products {
	return entities.ProductEntity(
		m.ID,
		m.ProductCode,
		m.ProductName,
		m.ProductPrice,
		m.CreatedBy,
		m.CreatedAt,
		m.UpdatedBy,
		m.UpdatedAt,
	)
}

func TransformProductModel(e entities.Products) *Product {
	return &Product{
		ID:           e.ID(),
		ProductCode:  e.ProductCode(),
		ProductName:  e.ProductName(),
		ProductPrice: e.ProductPrice(),
		CreatedBy:    e.CreatedBy(),
		CreatedAt:    e.CreatedAt(),
		UpdatedBy:    e.UpdatedBy(),
		UpdatedAt:    e.UpdatedAt(),
	}
}
