package entities

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	Id          uuid.UUID `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Permissions string    `json:"permissions" bson:"permissions"` // {"profile": "read;write;edit;delete", "question:" "generate;answer;edit;etc", "config": "read;write;" } format
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"deleted_at"`
}
