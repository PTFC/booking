package responses

import (
	"time"
)

type CustomerResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LicenseId   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Active      bool   `json:"active"`
}

type BookingResponse struct {
	ID         string    `json:"id"`
	CustomerId int64     `json:"customer_id"`
	FlightId   int64     `json:"flight_id"`
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	BookedDate time.Time `json:"booked_date"`
}
