package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

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
// func (r *AuthRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
//     err := r.users().FindOne(ctx, bson.M{"email": email}).Decode(&u)
// }
