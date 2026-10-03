package repositories

import (
	"context"
	"errors"

	"github.com/heronhoga/simplearn/backend/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const usersCollection = "users"

type AuthRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) users() *mongo.Collection {
	return r.db.Collection(usersCollection)
}

// usage
func (r *AuthRepository) FindExistingUserByEmail(context context.Context, email string) (*entities.User, error) {
	var existingUser entities.User
	filter := bson.M{"email": email}

	err := r.users().FindOne(context, filter).Decode(&existingUser)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &existingUser, nil
}

func (r *AuthRepository) InsertUser(context context.Context, newUser *entities.User) error {
	_, err := r.users().InsertOne(context, newUser)
	if err != nil {
		return err
	}

	return nil
}
