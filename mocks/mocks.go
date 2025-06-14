package mocks

import (
	"github.com/igormpb/travel-go/models"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (m *Repository) Create(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *Repository) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *Repository) GetByUserId(userID string) (*models.Credentials, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.Credentials), args.Error(1)
}

func (m *Repository) CreateCredentials(creds *models.Credentials) (*models.Credentials, error) {
	args := m.Called(creds)
	return args.Get(0).(*models.Credentials), args.Error(1)
}

func (m *Repository) CreateTravel(travel *models.Travel) (*models.Travel, error) {
	args := m.Called(travel)
	return args.Get(0).(*models.Travel), args.Error(1)
}

func (m *Repository) ListTravel(userID, status, destination, startDate, endDate string) ([]models.Travel, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Travel), args.Error(1)
}

func (m *Repository) GetTravelById(id, userID string) (*models.Travel, error) {
	args := m.Called(id, userID)
	return args.Get(0).(*models.Travel), args.Error(1)
}

func (m *Repository) UpdateTravel(travel *models.Travel) error {
	args := m.Called(travel)
	return args.Error(0)
}

func (m *Repository) CreateNotification(notification *models.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}

func (m *Repository) ListNotificationsByUser(userID string) ([]models.Notification, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Notification), args.Error(1)
}

func (m *Repository) MarkNotificationAsRead(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
