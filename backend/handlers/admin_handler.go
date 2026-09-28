package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/requests"
	"github.com/heronhoga/simplearn/backend/services"
	"github.com/heronhoga/simplearn/backend/utils"
)

type AdminHandler struct {
	service *services.AdminService
}

func NewAdminHandler(service *services.AdminService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

// module
func (h *AdminHandler) CreateModule(c fiber.Ctx) error {
	var createModuleRequest requests.CreateModuleRequest
	err := c.Bind().Body(&createModuleRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message":    "error",
			"error_type": "client",
			"error":      "invalid request body",
		})
	}

	errorTags, errorMessages, errorCode := utils.ValidateAndMapRequest(createModuleRequest)
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

	err = h.service.CreateModule(context.Background(), createModuleRequest)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message":    "error",
			"error_type": "server",
			"error":      "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}
