package factories

import (
	"ecommerce/helpers"
	"ecommerce/models/mysql"
)

func newProduct() *mysql.Product {
	return &mysql.Product{
		Name:             helpers.GenerateLetter(8),
		ChannelProductId: helpers.GenerateString(14),
		CommonColumn: mysql.CommonColumn{
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
	}
}

func DefinitionProduct() *mysql.Product {
	return newProduct()
}

func SeedProduct(number int) []*mysql.Product {
	data := make([]*mysql.Product, number)

	for i := 0; i < number; i++ {
		data[i] = DefinitionProduct()
	}

	return data
}
