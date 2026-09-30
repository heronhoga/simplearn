package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/handlers"
)

func InitAdminRoutes(app *fiber.App, adminHandler *handlers.AdminHandler) {
	// modules
	app.Post("/admin/modules/create", adminHandler.CreateModule)
	app.Put("/admin/modules/edit", adminHandler.EditModule)
	app.Put("/admin/modules/delete", adminHandler.DeleteModule)

	// permissions
	app.Post("/admin/permissions/create", adminHandler.CreatePermission)

}
