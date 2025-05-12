package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID string             `bson:"product_id" json:"product_id"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Status    string             `bson:"status" json:"status"` // created, paid, shipped
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
