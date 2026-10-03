package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/handlers"
	"github.com/heronhoga/simplearn/backend/middlewares"
)

func InitAdminRoutes(app *fiber.App, adminHandler *handlers.AdminHandler) {
	// modules
	app.Post("/admin/modules/create", middlewares.VerifyLoginStatus, adminHandler.CreateModule)
	app.Put("/admin/modules/edit", adminHandler.EditModule)
	app.Put("/admin/modules/delete", adminHandler.DeleteModule)

	// permissions
	app.Post("/admin/permissions/create", adminHandler.CreatePermission)

}
