package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Module struct {
	Id        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `json:"deleted_at" bson:"deleted_at"`
}
