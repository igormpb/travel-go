package models

import (
	"time"

	"github.com/google/uuid"
)

type TravelStatus string

const (
	StatusRequested TravelStatus = "REQUESTED"
	StatusApproved  TravelStatus = "APPROVED"
	StatusCanceled  TravelStatus = "CANCELED"
)

func IsValidStatus(s string) bool {
	switch TravelStatus(s) {
	case StatusRequested, StatusApproved, StatusCanceled:
		return true
	default:
		return false
	}
}

type Travel struct {
	ID            uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RequesterName string       `gorm:"not null" json:"requesterName"`
	Destination   string       `gorm:"not null" json:"destination"`
	DepartureDate time.Time    `gorm:"not null" json:"departureDate"`
	ReturnDate    time.Time    `gorm:"not null" json:"returnDate"`
	Status        TravelStatus `gorm:"type:varchar(20);default:'REQUESTED'" json:"status"`
	UserID        uuid.UUID    `gorm:"type:uuid;not null" json:"userId"`
	CreatedAt     time.Time    `gorm:"autoCreateTime" json:"createdAt"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}
