package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/heronhoga/simplearn/backend/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const modulesCollection = "modules"

type AdminRepository struct {
	db *mongo.Database
}

func NewAdminRepository(db *mongo.Database) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) modules() *mongo.Collection {
	return r.db.Collection(modulesCollection)
}

func (r *AdminRepository) InsertModule(context context.Context, module *entities.Module) error {
	_, err := r.modules().InsertOne(context, module)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminRepository) UpdateModule(ctx context.Context, id string, name string) error {
	update := bson.M{
		"$set": bson.M{
			"name":       name,
			"updated_at": time.Now(),
		},
	}

	filter := bson.M{
		"id": id,
	}

	result, err := r.modules().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("module with id %s not found", id)
	}

	return nil
}
