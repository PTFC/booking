package responses

import (
	"github.com/google/uuid"
)

type CustomerResponse struct {
	ID          uuid.UUID `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	LicenseId   string    `gorm:"column:license_id"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Email       int       `gorm:"column:email"`
	Password    int       `gorm:"column:password"`
	Active      int       `gorm:"column:active"`
}
