package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Question struct {
	Id            bson.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	DocumentId    bson.ObjectID     `json:"document_id" bson:"document_id"`
	UserId        bson.ObjectID     `json:"user_id" bson:"user_id"`
	Type          string            `json:"type" bson:"type"`
	Question      string            `json:"question" bson:"question"`
	Options       map[string]string `json:"options,omitempty" bson:"options,omitempty"`
	Answer        string            `json:"answer" bson:"answer"`
	Explanation   string            `json:"explanation" bson:"explanation"`
	SourceChunkId bson.ObjectID     `json:"source_chunk_id,omitempty" bson:"source_chunk_id,omitempty"`
	CreatedAt     time.Time         `json:"created_at" bson:"created_at"`
	DeletedAt     *time.Time        `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
