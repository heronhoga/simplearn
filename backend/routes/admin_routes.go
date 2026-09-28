package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/handlers"
)

func InitAdminRoutes(app *fiber.App, adminHandler *handlers.AdminHandler) {
	app.Post("/admin/modules/create", adminHandler.CreateModule)
}
