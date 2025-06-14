package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/igormpb/travel-go/app"
	"github.com/igormpb/travel-go/database"

	"github.com/igormpb/travel-go/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()

	db.AutoMigrate(
		&models.User{},
		&models.Credentials{},
		&models.Travel{},
		&models.Notification{},
	)

	appServer := fiber.New()
	appServer.Use(cors.New())

	app.App(appServer, db)

	log.Fatal(appServer.Listen(":3002"))
}
