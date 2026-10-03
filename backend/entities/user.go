package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id             bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email          string        `json:"email" bson:"email"`
	PasswordHashed string        `json:"password_hashed" bson:"password_hashed"`
	FullName       string        `json:"full_name" bson:"full_name"`
	Role           string        `json:"role" bson:"role"`
	Permissions    string        `json:"permissions" bson:"permissions"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
