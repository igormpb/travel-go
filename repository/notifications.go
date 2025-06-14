package repository

import (
	"github.com/igormpb/travel-go/models"
)

func (r *repository) CreateNotification(notification *models.Notification) error {
	return r.client.Create(notification).Error
}

func (r *repository) ListNotificationsByUser(userID string) ([]models.Notification, error) {
	var list []models.Notification
	err := r.client.Where("user_id = ?", userID).Order("created_at desc").Find(&list).Error
	return list, err
}

func (r *repository) MarkNotificationAsRead(id string) error {
	return r.client.Model(&models.Notification{}).
		Where("id = ?", id).
		Update("is_read", true).Error
}
