package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID `json:"id" bson:"id"`
	Email          string    `json:"email" bson:"email"`
	PasswordHashed string    `json:"password_hashed" bson:"password_hashed"`
	FullName       string    `json:"full_name" bson:"full_name"`
	Role           string    `json:"role" bson:"role"`
	Permissions    string    `json:"permissions" bson:"permissions"` // {"profile": "read;write;edit;delete", "question:" "generate;answer;edit;etc" } format
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at" bson:"deleted_at"`
}
