package responses

import (
	"time"

	"github.com/google/uuid"
)

type BookingResponse struct {
	ID         uuid.UUID `gorm:"column:id"`
	CustomerId int       `gorm:"column:customer_id"`
	FlightId   int       `gorm:"column:flight_id"`
	Code       string    `gorm:"column:code"`
	Status     string    `gorm:"column:status"`
	BookedDate int       `gorm:"column:booked_date"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
