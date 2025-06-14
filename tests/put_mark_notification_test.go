package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/travel-go/api"
	"github.com/igormpb/travel-go/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPutMarkNotification(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)

	id := uuid.New().String()
	app.Put("/notifications/:id/read", srv.MarkNotificationRead)

	mockRepo.On("MarkNotificationAsRead", id).Return(nil)

	req := httptest.NewRequest("PUT", "/notifications/"+id+"/read", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
