package entities

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	Id        uuid.UUID `json:"id"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
