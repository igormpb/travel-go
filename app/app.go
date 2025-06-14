package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igormpb/travel-go/api"
	"gorm.io/gorm"
)

func App(app *fiber.App, db *gorm.DB) {
	api.Api(app, db)
}
