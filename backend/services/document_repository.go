package services

import "github.com/heronhoga/simplearn/backend/repositories"

type DocumentService struct {
	repository *repositories.DocumentRepository
}

func NewDocumentService(repository *repositories.DocumentRepository) *DocumentService {
	return &DocumentService{
		repository: repository,
	}
}
