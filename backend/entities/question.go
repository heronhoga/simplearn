package entities

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id          uuid.UUID `json:"id"`
	DocumentId  uuid.UUID `json:"document_id"`
	UserId      uuid.UUID `json:"user_id"`
	Question    string    `json:"question"`
	Options     string    `json:"options"` // example: {"a": "asd", "b": "fgh"}
	Answer      string    `json:"answer"`  // example: "a" or "b"
	Explanation string    `json:"explanation"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
