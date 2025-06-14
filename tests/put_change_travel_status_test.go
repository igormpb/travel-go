package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/travel-go/api"
	"github.com/igormpb/travel-go/mocks"
	"github.com/igormpb/travel-go/models"
	"github.com/stretchr/testify/assert"
)

func TestPutChangeTravelStatus_OwnerBlocked(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)

	travelID := uuid.New().String()
	userID := uuid.New().String()
	app.Put("/travels/:id/status", func(c *fiber.Ctx) error {
		c.Locals("userId", userID)
		return srv.UpdateStatus(c)
	})

	mockRepo.On("GetTravelById", travelID, userID).Return(&models.Travel{
		ID:            uuid.MustParse(travelID),
		UserID:        uuid.MustParse(userID),
		Status:        models.StatusRequested,
		DepartureDate: time.Now().AddDate(0, 0, 10),
	}, nil)

	req := httptest.NewRequest("PUT", "/travels/"+travelID+"/status",
		bytes.NewBufferString(`{"status":"APPROVED"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, 403, resp.StatusCode)
}
