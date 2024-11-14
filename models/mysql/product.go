package mysql

type Product struct {
	CommonColumn
	Name             string `json:"name"`
	ChannelProductId string `json:"channel_product_id"`
}

func (Product) TableName() string {
	return "products"
}
