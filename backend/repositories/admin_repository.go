package repositories

import (
	"context"
	"errors"
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

func (r *AdminRepository) GetModuleByName(context context.Context, moduleName string) (*entities.Module, error) {
	var module entities.Module

	filter := bson.M{"name": moduleName}

	err := r.modules().FindOne(context, filter).Decode(&module)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &module, nil
}

func (r *AdminRepository) UpdateModule(ctx context.Context, id string, name string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid module id")
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "updated_at", Value: time.Now()},
		}},
	}

	result, err := r.modules().UpdateByID(ctx, objID, update)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("module not found")
	}

	return nil
}

func (r *AdminRepository) DeleteModule(ctx context.Context, moduleId string) error {
	objID, err := bson.ObjectIDFromHex(moduleId)
	if err != nil {
		return fmt.Errorf("invalid module id")
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deleted_at", Value: time.Now()},
		}},
	}

	result, err := r.modules().UpdateByID(ctx, objID, update)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("module not found")
	}

	return nil
}
