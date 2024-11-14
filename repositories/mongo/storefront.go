package mongo

import (
	"ecommerce/database/connection"
	"ecommerce/models/mongo"
)

type StorefrontRepository struct {
	BaseRepository[mongo.Storefront]
}

func NewStorefrontRepository(db *connection.MongoDatabase) *StorefrontRepository {
	return &StorefrontRepository{
		BaseRepository: *NewBaseRepository[mongo.Storefront](db, mongo.Storefront{}.CollectionName()),
	}
}
