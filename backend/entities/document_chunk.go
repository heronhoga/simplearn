package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DocumentChunk struct {
	Id         bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DocumentId bson.ObjectID `json:"document_id" bson:"document_id"`
	UserId     bson.ObjectID `json:"user_id" bson:"user_id"`

	Content    string `json:"content" bson:"content"`
	ChunkIndex int    `json:"chunk_index" bson:"chunk_index"`

	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
