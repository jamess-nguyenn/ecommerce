package mongo

import (
	"ecommerce/database/connection"
	"ecommerce/models/mongo"
)

type productRepository struct {
	baseRepository[mongo.Product]
}

func NewProductRepository(db *connection.MongoDatabase) *productRepository {
	return &productRepository{
		baseRepository: *newBaseRepository[mongo.Product](db, mongo.Product{}.CollectionName()),
	}
}
