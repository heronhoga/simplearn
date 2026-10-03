package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Document struct {
	Id        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FileName  string        `json:"file_name" bson:"file_name"`
	FilePath  string        `json:"file_path" bson:"file_path"`
	UserId    bson.ObjectID `json:"user_id" bson:"user_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	DeletedAt *time.Time    `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
