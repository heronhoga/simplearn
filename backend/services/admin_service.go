package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/heronhoga/simplearn/backend/entities"
	"github.com/heronhoga/simplearn/backend/repositories"
	"github.com/heronhoga/simplearn/backend/requests"
)

type AdminService struct {
	repository *repositories.AdminRepository
}

func NewAdminService(repository *repositories.AdminRepository) *AdminService {
	return &AdminService{
		repository: repository,
	}
}

func (s *AdminService) CreateModule(context context.Context, newModule requests.CreateModuleRequest) error {
	module := &entities.Module{
		Id:        uuid.New(),
		Name:      newModule.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repository.InsertModule(context, module)
	if err != nil {
		return err
	}

	return nil
}
