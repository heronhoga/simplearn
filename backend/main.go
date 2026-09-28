package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/simplearn/backend/config/database"
	"github.com/heronhoga/simplearn/backend/handlers"
	"github.com/heronhoga/simplearn/backend/repositories"
	"github.com/heronhoga/simplearn/backend/routes"
	"github.com/heronhoga/simplearn/backend/services"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system env instead")
	}

	// connect to mongodb
	client, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()

	dbName := os.Getenv("MONGODB_DATABASE")
	db := client.Database(dbName)

	// Set Repository
	authRepository := repositories.NewAuthRepository(db)

	// Set Services
	authService := services.NewAuthService(authRepository)

	// Set Handlers
	authHandler := handlers.NewAuthHandler(authService)

	app := fiber.New()

	// Set Routes
	routes.InitAuthRoutes(app, authHandler)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal("Error starting simplearn backend service")
	}
}
