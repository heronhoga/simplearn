package entities

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id        string    `json:"id" bson:"id"`
	FileName  string    `json:"file_name" bson:"file_name"`
	FilePath  string    `json:"file_path" bson:"file_path"`
	UserId    uuid.UUID `json:"user_id" bson:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}
