package transforms

import "github.com/milfan/go-boilerplate/internal/api/entities"

type ProductList struct {
	ID           uint64  `json:"id"`
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	ProductPrice float64 `json:"productPrice"`
}

func TransformProductList(entity []entities.Products) []ProductList {
	data := make([]ProductList, 0)
	for _, item := range entity {
		data = append(data, ProductList{
			ID:           item.ID(),
			ProductCode:  item.ProductCode(),
			ProductName:  item.ProductName(),
			ProductPrice: item.ProductPrice(),
		})
	}
	return data
}
