package mysql

import (
	"ecommerce/database/connection"
	"ecommerce/models/mysql"
)

type storefrontRepository struct {
	baseRepository[mysql.Storefront]
}

func NewStorefrontRepository(db *connection.MysqlDatabase) *storefrontRepository {
	return &storefrontRepository{
		baseRepository: *newBaseRepository[mysql.Storefront](db),
	}
}
