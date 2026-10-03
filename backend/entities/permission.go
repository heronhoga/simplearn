package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Permission struct {
	Id        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Action    string        `json:"action" bson:"action"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at" bson:"deleted_at"`
}
