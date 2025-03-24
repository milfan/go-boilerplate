package models

import (
	"time"

	"github.com/milfan/go-boilerplate/internal/api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	ID           uint64         `gorm:"column:id;primaryKey"`
	OrderCode    string         `gorm:"column:order_code"`
	OrderDate    time.Time      `gorm:"column:order_date"`
	CreatedBy    string         `gorm:"column:created_by"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedBy    string         `gorm:"column:updated_by"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	OrderDetails []OrderDetail  `gorm:"foreignKey:OrderID;references:ID"`
}

func (m Order) BeforeSave(tx *gorm.DB) error {
	tx.Statement.AddClause(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		UpdateAll: true,
	})
	return nil
}

func (m Order) Entity() *entities.Orders {

	orderDetails := make([]entities.OrderDetails, 0)
	for _, item := range m.OrderDetails {
		t := item.Entity()
		orderDetails = append(orderDetails, *t)
	}

	return entities.OrderEntity(
		m.ID,
		m.OrderCode,
		m.OrderDate,
		m.CreatedBy,
		m.CreatedAt,
		m.UpdatedBy,
		m.UpdatedAt,
		orderDetails,
	)
}

func TransformOrderModel(e entities.Orders) *Order {

	orderDetails := make([]OrderDetail, 0)
	for _, item := range e.OrderDetails() {
		m := TransformOrderDetailModel(item)
		orderDetails = append(orderDetails, *m)
	}

	return &Order{
		ID:           e.ID(),
		OrderCode:    e.OrderCode(),
		OrderDate:    e.OrderDate(),
		CreatedBy:    e.CreatedBy(),
		CreatedAt:    e.CreatedAt(),
		UpdatedBy:    e.UpdatedBy(),
		UpdatedAt:    e.UpdatedAt(),
		OrderDetails: orderDetails,
	}
}
