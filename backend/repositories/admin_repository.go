package repositories

import (
	"context"

	"github.com/heronhoga/simplearn/backend/entities"
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
