package mysql

type Product struct {
	CommonColumn
	Name             string `json:"name" gorm:"column:name"`
	ChannelProductId string `json:"channel_product_id" gorm:"column:channel_product_id"`
}

func (Product) TableName() string {
	return "products"
}
