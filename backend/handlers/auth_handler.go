package handlers

import (
	"github.com/gofiber/fiber/v3"
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

func (h *AuthHandler) Register(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}
