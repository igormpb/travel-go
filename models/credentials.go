package models

import "github.com/google/uuid"

type Credentials struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"userId"`
	Salt         string    `gorm:"not null" json:"salt"`
	PasswordHash string    `gorm:"not null" json:"passwordHash"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}
