package mongo

import (
	"ecommerce/database/connection"
	"ecommerce/models/mongo"
)

type storefrontRepository struct {
	baseRepository[mongo.Storefront]
}

func NewStorefrontRepository(db *connection.MongoDatabase) *storefrontRepository {
	return &storefrontRepository{
		baseRepository: *newBaseRepository[mongo.Storefront](db, mongo.Storefront{}.CollectionName()),
	}
}
