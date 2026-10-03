package services

import (
	"context"
	"errors"
	"time"

	"github.com/heronhoga/simplearn/backend/entities"
	"github.com/heronhoga/simplearn/backend/repositories"
	"github.com/heronhoga/simplearn/backend/requests"
	"github.com/heronhoga/simplearn/backend/responses"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *repositories.AuthRepository
}

func NewAuthService(repository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Register(context context.Context, newUser requests.RegisterRequest) (responses.RegisterResponse, error) {
	// find existing user
	existingUser, err := s.repository.FindExistingUserByEmail(context, newUser.Email)
	if err != nil {
		return responses.RegisterResponse{}, errors.New("internal server error") // empty struct and err
	}

	if existingUser != nil {
		return responses.RegisterResponse{}, errors.New("user already exists")
	}

	// insert new user
	// hash password first
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.RegisterResponse{}, errors.New("internal server error")
	}
	hashedPasswordString := string(bytes)

	newUserInput := &entities.User{
		Email:          newUser.Email,
		PasswordHashed: hashedPasswordString,
		FullName:       newUser.FullName,
		Role:           "user",
		Permissions:    `{"profile": "read;write;edit;delete", "question": "generate;answer;delete" }`,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = s.repository.InsertUser(context, newUserInput)
	if err != nil {
		return responses.RegisterResponse{}, errors.New("internal server error")
	}

	return responses.RegisterResponse{Message: "register successful, please login using the credentials"}, nil

}
