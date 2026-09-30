package services

import (
	"context"
	"errors"
	"fmt"
	"time"

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

// string: message; error: system's error
func (s *AdminService) CreateModule(context context.Context, newModule requests.CreateModuleRequest) (string, error) {
	module := &entities.Module{
		Name:      newModule.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// find existing module
	existingModule, err := s.repository.GetModuleByName(context, module.Name)
	if err != nil {
		fmt.Println(err)
		return "internal server error", err
	}

	if existingModule != nil {
		return "module already exists", errors.New("module already exists")
	}

	err = s.repository.InsertModule(context, module)
	if err != nil {
		return "internal server error", err
	}

	return "ok", nil
}

func (s *AdminService) EditModule(context context.Context, newModule requests.EditModuleRequest) error {
	err := s.repository.UpdateModule(context, newModule.Id, newModule.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminService) DeleteModule(context context.Context, moduleId string) error {
	err := s.repository.DeleteModule(context, moduleId)
	if err != nil {
		return err
	}

	return nil
}
