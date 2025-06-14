package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igormpb/travel-go/middleware"
	"github.com/igormpb/travel-go/repository"
	"gorm.io/gorm"
)

type Service struct {
	repository repository.IRepository
}

func NewService(repo repository.IRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func Api(app *fiber.App, db *gorm.DB) {
	repository := repository.NewRepository(db)
	srv := NewService(repository)

	//Auth
	auth := app.Group("/auth")

	auth.Post("/singup", srv.PostSignup)
	auth.Post("/signin", srv.PostSignin)

	auth.Post("/travel", middleware.JWT(), srv.CreateTravel)

	api := app.Group("/api/v1", middleware.JWT())
	api.Post("/travels", srv.CreateTravel)
	api.Get("/travels", srv.ListTravels)
	api.Get("/travels/:id", srv.GetTravelByID)
	api.Put("/travels/:id/status", srv.UpdateStatus)

	api.Get("/notifications", srv.ListNotifications)
	api.Put("/notifications/:id/read", srv.MarkNotificationRead)
}
