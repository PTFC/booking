package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()`
	CustomerId int64     `gorm:"type:integer"`
	FlightId   int64     `gorm:"type:integer"`
	Code       string    `gorm:"type:varchar(256)"`
	Status     string    `gorm:"type:varchar(256)"`
	BookedDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
