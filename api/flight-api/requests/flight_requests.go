package requests

import (
	"time"
)

type CreateFlightRequest struct {
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	AvailableSlot int64     `json:"available_slot" binding:"required"`
}

type UpdateFlightRequest struct {
	Id            string    `json:"id" binding:"required"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"available_slot"`
}

type SearchFlightRequest struct {
	QueryString string `json:"q" binding:"required"`
}
