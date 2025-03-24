package entities

import (
	"time"
)

type OrderDetails struct {
	BaseEntity
	id           uint64
	orderId      uint64
	productId    uint64
	productQty   uint32
	productPrice float64
	subtotal     float64
	product      *Products
}

func (e *OrderDetails) ID() uint64 {
	return e.id
}

func (e *OrderDetails) OrderID() uint64 {
	return e.orderId
}

func (e *OrderDetails) ProductID() uint64 {
	return e.productId
}

func (e *OrderDetails) Qty() uint32 {
	return e.productQty
}

func (e *OrderDetails) Price() float64 {
	return e.productPrice
}

func (e *OrderDetails) Subtotal() float64 {
	return float64(e.Qty()) * e.Price()
}

func (e *OrderDetails) Product() *Products {
	return e.product
}

func (e *OrderDetails) SetProduct(val *Products) {
	e.product = val
}

func OrderDetailEntity(
	id uint64,
	orderId, productId uint64,
	productQty uint32,
	productPrice float64,
	createdBy string,
	createdAt time.Time,
	updatedBy string,
	updatedAt time.Time,
) *OrderDetails {
	return &OrderDetails{
		BaseEntity: BaseEntity{
			createdBy: createdBy,
			createdAt: createdAt,
			updatedBy: updatedBy,
			updatedAt: updatedAt,
		},
		id:           id,
		orderId:      orderId,
		productId:    productId,
		productQty:   productQty,
		productPrice: productPrice,
	}
}

func NewOrderDetail(
	productId uint64,
	productQty uint32,
	productPrice float64,
	createdBy string,
) *OrderDetails {
	return &OrderDetails{
		BaseEntity:   New(createdBy),
		productId:    productId,
		productQty:   productQty,
		productPrice: productPrice,
	}
}
