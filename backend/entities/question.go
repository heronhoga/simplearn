package entities

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id          uuid.UUID `json:"id" bson:"id"`
	DocumentId  uuid.UUID `json:"document_id" bson:"document_id"`
	UserId      uuid.UUID `json:"user_id" bson:"user_id"`
	Question    string    `json:"question" bson:"question"`
	Options     string    `json:"options" bson:"options"` // example: {"a": "asd", "b": "fgh"}
	Answer      string    `json:"answer" bson:"answer"`   // example: "a" or "b"
	Explanation string    `json:"explanation" bson:"explanation"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"created_at"`
}
