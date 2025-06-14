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

func TestListNotifications(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)

	app.Get("/notifications", func(c *fiber.Ctx) error {
		c.Locals("userId", uuid.New().String())
		return srv.ListNotifications(c)
	})

	mockRepo.On("ListNotificationsByUser", mock.Anything).
		Return([]models.Notification{{Message: "Status changed"}}, nil)

	req := httptest.NewRequest("GET", "/notifications", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
