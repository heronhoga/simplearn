package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Role struct {
	Id          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Permissions string        `json:"permissions" bson:"permissions"` // {"profile": "read;write;edit;delete", "question:" "generate;answer;edit;etc", "config": "read;write;" } format
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at" bson:"deleted_at"`
}
