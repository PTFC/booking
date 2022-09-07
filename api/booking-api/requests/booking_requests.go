package requests

import (
	"time"
)

type CreateBookingRequest struct {
	CustomerId int64     `json:"customer_id" binding:"required"`
	FlightId   int64     `json:"flight_id" binding:"required"`
	Code       string    `json:"code" binding:"required"`
	Status     string    `json:"status" binding:"required"`
	BookedDate time.Time `json:"booked_date" binding:"required"`
}

type FindBookingRequest struct {
	Id string `json:"id" binding:"required"`
}

type CancelBookingRequest struct {
	Id string `json:"id" binding:"required"`
}
