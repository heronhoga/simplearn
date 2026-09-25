package entities

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id        uuid.UUID `json:"id"`
	FileName  string    `json:"file_name"`
	FilePath  string    `json:"file_path"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
