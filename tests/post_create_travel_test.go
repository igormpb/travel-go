package tests

import (
	"bytes"
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

func TestPostCreateTravel_Success(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)
	app.Post("/travels", func(c *fiber.Ctx) error {
		c.Locals("userId", uuid.New().String())
		return srv.CreateTravel(c)
	})

	payload := `{
	"requesterName": "Igor",
	"destination": "Rio",
	"departureDate": "2025-07-10",
	"returnDate": "2025-07-15"
}`
	req := httptest.NewRequest("POST", "/travels", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	mockRepo.On("CreateTravel", mock.Anything).Return(&models.Travel{ID: uuid.New(), RequesterName: "Igor"}, nil)

	resp, _ := app.Test(req)
	assert.Equal(t, 201, resp.StatusCode)
}
