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
	"github.com/stretchr/testify/mock"
)

func TestGetTravels(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)
	app.Get("/travels", func(c *fiber.Ctx) error {
		c.Locals("userId", uuid.New().String())
		return srv.ListTravels(c)
	})

	mockTravels := []models.Travel{{ID: uuid.New(), RequesterName: "X"}}
	mockRepo.On("ListTravel", mock.Anything).Return(mockTravels, nil)

	req := httptest.NewRequest("GET", "/travels", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
