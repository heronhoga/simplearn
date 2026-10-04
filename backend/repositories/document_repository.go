package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

const documentsCollection = "documents"

type DocumentRepository struct {
	db *mongo.Database
}

func NewDocumentRepository(db *mongo.Database) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *AuthRepository) documents() *mongo.Collection {
	return r.db.Collection(documentsCollection)
}
