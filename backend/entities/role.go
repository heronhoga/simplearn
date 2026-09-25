package entities

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Permissions string    `json:"permissions"` // {"profile": "read;write;edit;delete", "question:" "generate;answer;edit;etc", "config": "read;write;" } format
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
