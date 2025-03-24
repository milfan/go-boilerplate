package entities

import (
	"strconv"
	"time"
)

type Products struct {
	BaseEntity
	id           uint64
	productCode  string
	productName  string
	productPrice float64
}

func (e *Products) ID() uint64 {
	return e.id
}

func (e *Products) ProductCode() string {
	return e.productCode
}

func (e *Products) ProductName() string {
	return e.productName
}

func (e *Products) ProductPrice() float64 {
	return e.productPrice
}

func (e *Products) GenerateProductCode(lastId uint64) string {
	code := "PRODUCT-"
	year, month, _ := time.Now().Date()

	s := strconv.FormatUint(uint64(year+99), 10)
	code = code + s
	u := strconv.FormatUint(uint64(month+12), 10)
	code = code + u
	o := strconv.FormatUint(lastId, 10)
	code = code + o
	return code
}

func ProductEntity(
	id uint64,
	productCode, productName string,
	productPrice float64,
	createdBy string,
	createdAt time.Time,
	updatedBy string,
	updatedAt time.Time,
) *Products {
	return &Products{
		BaseEntity: BaseEntity{
			createdBy: createdBy,
			createdAt: createdAt,
			updatedBy: updatedBy,
			updatedAt: updatedAt,
		},
		id:           id,
		productCode:  productCode,
		productName:  productName,
		productPrice: productPrice,
	}
}

func NewProduct(
	productName string,
	productPrice float64,
	createdBy string,
) *Products {
	return &Products{
		BaseEntity:   New(createdBy),
		productName:  productName,
		productPrice: productPrice,
	}
}
