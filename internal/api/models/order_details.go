package models

import (
	"time"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID           uint64         `gorm:"column:id;primaryKey"`
	OrderID      uint64         `gorm:"column:order_id"`
	ProductID    uint64         `gorm:"column:product_id"`
	ProductQty   uint32         `gorm:"column:product_qty"`
	ProductPrice float64        `gorm:"column:product_price"`
	CreatedBy    string         `gorm:"column:created_by"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedBy    string         `gorm:"column:updated_by"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Product      *Product       `gorm:"foreignKey:ProductID;references:ID"`
}

func (m OrderDetail) Entity() *entities.OrderDetails {

	entity := entities.OrderDetailEntity(
		m.ID,
		m.OrderID,
		m.ProductID,
		m.ProductQty,
		m.ProductPrice,
		m.CreatedBy,
		m.CreatedAt,
		m.UpdatedBy,
		m.UpdatedAt,
	)

	if m.Product != nil {
		productEntity := m.Product.Entity()
		entity.SetProduct(productEntity)
	}
	return entity
}

func TransformOrderDetailModel(e entities.OrderDetails) *OrderDetail {
	return &OrderDetail{
		ID:           e.ID(),
		OrderID:      e.OrderID(),
		ProductID:    e.ProductID(),
		ProductQty:   e.Qty(),
		ProductPrice: e.Price(),
		CreatedBy:    e.CreatedBy(),
		CreatedAt:    e.CreatedAt(),
		UpdatedBy:    e.UpdatedBy(),
		UpdatedAt:    e.UpdatedAt(),
	}
}
