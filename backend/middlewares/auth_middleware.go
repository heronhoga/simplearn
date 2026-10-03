package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/utils"
)

func VerifyLoginStatus(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid authorization header",
		})
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := utils.VerifyToken(accessToken)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid or expired token",
		})
	}

	return c.Next()
}
