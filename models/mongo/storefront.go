package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// maybe apply this for catefory
type Storefront struct {
	Id                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MarketplaceStorefrontId string             `bson:"marketplace_storefront_id,omitempty" json:"marketplace_storefront_id"`
	IsImport                bool               `bson:"is_import,omitempty" json:"is_import"`
	Data                    string             `bson:"data" json:"data"`
	CreatedAt               *time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt               *time.Time         `bson:"updated_at" json:"updated_at"`
}

func (Storefront) CollectionName() string {
	return "storefronts"
}
