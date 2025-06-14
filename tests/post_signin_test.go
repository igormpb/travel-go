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
	"golang.org/x/crypto/bcrypt"
)

func TestPostSignin_InvalidCreds(t *testing.T) {
	app := fiber.New()
	mockRepo := new(mocks.Repository)
	srv := api.NewService(mockRepo)
	app.Post("/auth/signin", srv.PostSignin)
	userID := uuid.New()

	hashed, _ := bcrypt.GenerateFromPassword([]byte("wrong"+"salt"), bcrypt.DefaultCost)

	user := models.User{ID: userID, Email: "u@u.com", Name: "U"}
	creds := models.Credentials{PasswordHash: string(hashed), Salt: "salt"}

	mockRepo.On("GetUserByEmail", "u@u.com").Return(&user, nil)
	mockRepo.On("GetByUserId", userID.String()).Return(&creds, nil)

	req := httptest.NewRequest("POST", "/auth/signin", bytes.NewBufferString(`{"email":"u@u.com","password":"1200"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, 401, resp.StatusCode)
}
