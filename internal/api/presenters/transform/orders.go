package transforms

import "github.com/milfan/go-boilerplate/internal/api/entities"

type OrderList struct {
	ID           uint64            `json:"id"`
	OrderDate    string            `json:"orderDate"`
	OrderTotal   float64           `json:"orderTotal"`
	OrderDetails []OrderDetailList `json:"orderdetail"`
}

type OrderDetailList struct {
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	ProductQty   uint32  `json:"productQty"`
	ProductPrice float64 `json:"productPrice"`
	Subtotal     float64 `json:"subtotal"`
}

func TransformOrderList(entity []entities.Orders) []OrderList {
	data := make([]OrderList, 0)
	for _, item := range entity {

		orderDetails := make([]OrderDetailList, 0)
		for _, itemDetail := range item.OrderDetails() {

			product := itemDetail.Product()
			orderDetail := OrderDetailList{
				ProductQty:   itemDetail.Qty(),
				ProductPrice: itemDetail.Price(),
				Subtotal:     itemDetail.Subtotal(),
			}
			if product != nil {
				orderDetail.ProductCode = product.ProductCode()
				orderDetail.ProductName = product.ProductName()
			}
			orderDetails = append(orderDetails, orderDetail)
		}

		data = append(data, OrderList{
			ID:           item.ID(),
			OrderDate:    item.CreatedAtAsISOString(),
			OrderTotal:   item.OrderTotal(),
			OrderDetails: orderDetails,
		})
	}
	return data
}
