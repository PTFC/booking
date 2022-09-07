package responses

import (
	"time"

	"github.com/google/uuid"
)

type FlightResponse struct {
	ID            uuid.UUID `gorm:"column:id"`
	Name          string    `gorm:"column:name"`
	From          string    `gorm:"column:from"`
	To            string    `gorm:"column:to"`
	Date          int       `gorm:"column:date"`
	Status        int       `gorm:"column:status"`
	AvailableSlot int       `gorm:"column:available_slot"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}
