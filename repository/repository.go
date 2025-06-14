package repository

import (
	"github.com/igormpb/travel-go/models"
	"gorm.io/gorm"
)

type INotificationsRepository interface {
	CreateNotification(notification *models.Notification) error
	ListNotificationsByUser(userID string) ([]models.Notification, error)
	MarkNotificationAsRead(id string) error
}
type ITravelRepository interface {
	CreateTravel(travel *models.Travel) (*models.Travel, error)
	ListTravel(userID, status, destination, startDate, endDate string) ([]models.Travel, error)
	GetTravelById(id, userID string) (*models.Travel, error)
	UpdateTravel(travel *models.Travel) error
}

type IUser interface {
	Create(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type ICredentials interface {
	GetByUserId(userID string) (*models.Credentials, error)
	CreateCredentials(creds *models.Credentials) (*models.Credentials, error)
}

type IRepository interface {
	ITravelRepository
	IUser
	INotificationsRepository
	ICredentials
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) IRepository {
	return &repository{client: client}
}
