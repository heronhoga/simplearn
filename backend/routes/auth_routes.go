package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/handlers"
)

func InitAuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	app.Post("/auth/login", authHandler.Login)
}
