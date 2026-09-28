package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/requests"
	"github.com/heronhoga/simplearn/backend/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var loginRequest requests.LoginRequest
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := c.Bind().Body(&loginRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error",
			"error":   "invalid request",
		})
	}

	err = validate.Struct(loginRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
		"data":    loginRequest,
	})
}
