package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	PasswordHashed string    `json:"password_hashed"`
	FullName       string    `json:"full_name"`
	Role           string    `json:"role"`
	Permissions    string    `json:"permissions"` // {"profile": "read;write;edit;delete", "question:" "generate;answer;edit;etc" } format
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
