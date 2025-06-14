package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/igormpb/travel-go/api"
	"github.com/igormpb/travel-go/mocks"
	"github.com/igormpb/travel-go/models"
	"github.com/stretchr/testify/assert"
)

func TestPostSignup_EmailExists(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)
	app.Post("/auth/signup", srv.PostSignup)

	mockRepo.On("GetUserByEmail", "test@example.com").Return(&models.User{}, nil)

	req := httptest.NewRequest("POST", "/auth/signup",
		bytes.NewBufferString(`{"name":"Test","email":"test@example.com","password":"123"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, 409, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}
