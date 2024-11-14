package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name             string             `bson:"name" json:"name"`
	ChannelProductId string             `bson:"channel_product_id,omitempty" json:"channel_product_id"`
	Data             string             `bson:"data" json:"data"`
	ExtraData        string             `bson:"extra_data" json:"extra_data"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}

func (Product) CollectionName() string {
	return "products"
}
