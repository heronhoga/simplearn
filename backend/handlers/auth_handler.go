package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/requests"
	"github.com/heronhoga/simplearn/backend/services"
	"github.com/heronhoga/simplearn/backend/utils"
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
	err := c.Bind().Body(&loginRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message":    "error",
			"error_type": "client",
			"error":      "invalid request body",
		})
	}

	errorTags, errorMessages, errorCode := utils.ValidateAndMapRequest(loginRequest)
	switch errorCode {
	case 500:
		return c.Status(500).JSON(fiber.Map{
			"message":    "error",
			"error_type": "server",
			"error":      "internal server error",
		})
	case 400:
		return c.Status(400).JSON(fiber.Map{
			"message":        "error",
			"error_type":     "validator",
			"error_tags":     errorTags,
			"error_messages": errorMessages,
		})
	}

	// hit service layer

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
		"data":    loginRequest,
	})
}
