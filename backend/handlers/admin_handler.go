package handlers

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/requests"
	"github.com/heronhoga/simplearn/backend/services"
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
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := c.Bind().Body(&createModuleRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message":    "error",
			"error_type": "client",
			"error":      "invalid request body",
		})
	}

	err = validate.Struct(createModuleRequest)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errorTags := []string{}
		errorMessages := []string{}
		if errors.As(err, &validationErrors) {

			for _, fieldErr := range validationErrors {
				var errorMessage string
				if fieldErr.Tag() == "required" {
					errorMessage = fieldErr.Field() + " is required"
				}
				errorTags = append(errorTags, fieldErr.Field())
				errorMessages = append(errorMessages, errorMessage)
			}

			return c.Status(400).JSON(fiber.Map{
				"message":       "error",
				"error_type":    "validator",
				"error":         errorTags,
				"error_message": errorMessages,
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"message":    "error",
				"error_type": "server",
				"error":      "internal server error",
			})
		}
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
