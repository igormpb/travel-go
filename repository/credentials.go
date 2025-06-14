package repository

import (
	"errors"

	"github.com/igormpb/travel-go/models"
	"gorm.io/gorm"
)

func (r *repository) GetByUserId(userID string) (*models.Credentials, error) {
	var creds models.Credentials
	if err := r.client.Where("user_id = ?", userID).First(&creds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &creds, nil
}

func (r *repository) CreateCredentials(creds *models.Credentials) (*models.Credentials, error) {
	if err := r.client.Create(creds).Error; err != nil {
		return nil, err
	}
	return creds, nil
}
