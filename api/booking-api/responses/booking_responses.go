package responses

import (
	"time"
)

type BookingResponse struct {
	ID         string    `json:"id"`
	CustomerId int64     `json:"customer_id"`
	FlightId   int64     `json:"flight_id"`
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	BookedDate time.Time `json:"booked_date"`
}
