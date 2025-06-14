package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/travel-go/api"
	"github.com/igormpb/travel-go/mocks"
	"github.com/igormpb/travel-go/models"
	"github.com/stretchr/testify/assert"
)

func TestGetTravelByID(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)

	travelID := uuid.New().String()
	userID := uuid.New().String()

	app.Get("/travels/:id", func(c *fiber.Ctx) error {
		c.Locals("userId", userID)
		return srv.GetTravelByID(c)
	})

	mockRepo.On("GetTravelById", travelID, userID).Return(&models.Travel{ID: uuid.MustParse(travelID)}, nil)

	req := httptest.NewRequest("GET", "/travels/"+travelID, nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
