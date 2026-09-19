package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	if err := app.Listen(":8000"); err != nil {
		log.Fatal("Error starting simplearn backend service")
	}
}
