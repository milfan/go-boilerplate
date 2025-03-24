package entities

import (
	"strconv"
	"time"
)

type Orders struct {
	BaseEntity
	id           uint64
	orderCode    string
	orderDate    time.Time
	orderTotal   float64
	orderDetails []OrderDetails
}

func (e *Orders) ID() uint64 {
	return e.id
}

func (e *Orders) OrderCode() string {
	return e.orderCode
}

func (e *Orders) OrderDate() time.Time {
	return e.orderDate
}

func (e *Orders) OrderTotal() float64 {
	return e.orderTotal
}

func (e *Orders) OrderDetails() []OrderDetails {
	return e.orderDetails
}

func (e *Orders) GenerateOrderCode(lastId uint64) string {
	code := "ORDER-"
	year, month, _ := time.Now().Date()

	s := strconv.FormatUint(uint64(year+99), 10)
	code = code + s
	u := strconv.FormatUint(uint64(month+12), 10)
	code = code + u
	o := strconv.FormatUint(lastId, 10)
	code = code + o
	return code
}

func OrderEntity(
	id uint64,
	orderCode string,
	orderDate time.Time,
	createdBy string,
	createdAt time.Time,
	updatedBy string,
	updatedAt time.Time,
	orderDetails []OrderDetails,
) *Orders {

	var orderTotal float64
	for _, orderDetail := range orderDetails {
		total := float64(orderDetail.Qty()) * orderDetail.Price()
		orderTotal += total
	}

	return &Orders{
		BaseEntity: BaseEntity{
			createdBy: createdBy,
			createdAt: createdAt,
			updatedBy: updatedBy,
			updatedAt: updatedAt,
		},
		id:           id,
		orderCode:    orderCode,
		orderDate:    orderDate,
		orderTotal:   orderTotal,
		orderDetails: orderDetails,
	}
}

func NewOrder(
	orderDate time.Time,
	createdBy string,
	orderDetails []OrderDetails,
) *Orders {
	return &Orders{
		BaseEntity:   New(createdBy),
		orderDate:    orderDate,
		orderDetails: orderDetails,
	}
}
