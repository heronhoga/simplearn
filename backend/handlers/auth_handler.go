package handlers

import (
	"context"

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
	loginResponse, err := h.service.Login(context.Background(), loginRequest)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":    "error",
			"error_type": "server",
			"error":      err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
		"data":    loginResponse,
	})
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var registerRequest requests.RegisterRequest
	if err := c.Bind().Body(&registerRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message":    "error",
			"error_type": "client",
			"error":      "invalid request body",
		})
	}

	errorTags, errorMessages, errorCode := utils.ValidateAndMapRequest(registerRequest)
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
	registerResponse, err := h.service.Register(context.Background(), registerRequest)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":    "error",
			"error_type": "server",
			"error":      err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": registerResponse.Message,
	})
}
