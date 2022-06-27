package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Product is the model that governs all notes objects retrived or inserted into the DB
type Product struct {
	ID           primitive.ObjectID `json:"_id"`
	Product_id   string             `json:"product_ID"`
	Product_type *string            `json:"product_type"`
	Product_name *string            `json:"product_name"`
	Size         *string            `json:"size"`
	Colour       *string            `json:"colour"`
	Price        *string            `json:"price"`
	Quantity     *string            `json:"quantity"`
	Description  *string            `json:"description"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
}
