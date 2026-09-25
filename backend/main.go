package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/config/database"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system env instead")
	}

	// connect to mongodb
	_, err = database.ConnectMongoDB()
	if err != nil {
		log.Fatalln(err)
	}

	// Set Repository
	// Set Services
	// Set Layers

	app := fiber.New()
	if err := app.Listen(":8000"); err != nil {
		log.Fatal("Error starting simplearn backend service")
	}
}
