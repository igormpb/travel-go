package repository

import (
	"errors"

	"github.com/igormpb/travel-go/models"
	"gorm.io/gorm"
)

func (r *repository) Create(user *models.User) (*models.User, error) {
	if err := r.client.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.client.Preload("Credentials").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
