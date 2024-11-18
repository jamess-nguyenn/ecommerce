package mysql

import (
	"ecommerce/database/connection"
	"ecommerce/models/mysql"
)

type productRepository struct {
	baseRepository[mysql.Product]
}

func NewProductRepository(db *connection.MysqlDatabase) *productRepository {
	return &productRepository{
		baseRepository: *newBaseRepository[mysql.Product](db),
	}
}
