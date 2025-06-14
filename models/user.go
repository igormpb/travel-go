package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Email       string       `gorm:"uniqueIndex;not null" json:"email"`
	Credentials *Credentials `gorm:"constraint:OnDelete:CASCADE" json:"credentials,omitempty"`
	Travels     []Travel     `gorm:"constraint:OnDelete:CASCADE" json:"travels,omitempty"`
}
