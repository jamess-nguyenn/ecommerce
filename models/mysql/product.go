package mysql

type Product struct {
	CommonColumn
	ProductName          string `json:"product_name" gorm:"column:name"`
	MarketplaceProductId string `json:"marketplace_product_id" gorm:"column:marketplace_product_id"`
}

func (Product) TableName() string {
	return "products"
}
