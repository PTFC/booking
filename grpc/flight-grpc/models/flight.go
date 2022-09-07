package models

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()`
	Name          string    `gorm:"type:varchar(256)"`
	From          string    `gorm:"type:varchar(256)"`
	To            string    `gorm:"type:varchar(256);not null;unquie"`
	Date          time.Time
	Status        string `gorm:"type:varchar(256)"`
	AvailableSlot int64  `gorm:"type:integer"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
