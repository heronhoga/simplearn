package entities

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	Id        uuid.UUID `json:"id" bson:"id"`
	Action    string    `json:"action" bson:"action"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}
