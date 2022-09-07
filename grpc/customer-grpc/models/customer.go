package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()`
	Name        string    `gorm:"type:varchar(256)"`
	Address     string    `gorm:"type:varchar(256)"`
	LicenseId   string    `gorm:"type:varchar(256)"`
	PhoneNumber string    `gorm:"type:varchar(256)"`
	Email       string    `gorm:"type:varchar(256)"`
	Password    string    `gorm:"type:varchar(256)"`
	Active      bool      `gorm:"type:bool"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
