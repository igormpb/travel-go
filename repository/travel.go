package repository

import (
	"errors"

	"github.com/igormpb/travel-go/models"
	"gorm.io/gorm"
)

func (r *repository) CreateTravel(travel *models.Travel) (*models.Travel, error) {
	if err := r.client.Create(travel).Error; err != nil {
		return nil, err
	}
	return travel, nil
}

func (r *repository) ListTravel(userID, status, destination, startDate, endDate string) ([]models.Travel, error) {
	var travels []models.Travel

	query := r.client.Preload("User")

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if destination != "" {
		query = query.Where("destination ILIKE ?", "%"+destination+"%")
	}
	if startDate != "" && endDate != "" {
		// Considerando que você está filtrando pela data de ida
		query = query.Where("departure_date BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Order("created_at desc").Find(&travels).Error
	if err != nil {
		return nil, err
	}

	return travels, nil
}

func (r *repository) GetTravelById(id, userID string) (*models.Travel, error) {
	var travel models.Travel
	if err := r.client.Preload("User").Where("id = ?", id).First(&travel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &travel, nil
}

func (r *repository) UpdateTravel(travel *models.Travel) error {
	return r.client.Save(travel).Error
}
