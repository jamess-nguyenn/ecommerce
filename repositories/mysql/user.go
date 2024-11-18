package mysql

import (
	"ecommerce/database/connection"
	"ecommerce/models/mysql"
)

type userRepository struct {
	baseRepository[mysql.User]
}

func NewUserRepository(db *connection.MysqlDatabase) *userRepository {
	return &userRepository{
		baseRepository: *newBaseRepository[mysql.User](db),
	}
}
