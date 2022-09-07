package requests

import "github.com/google/uuid"

type CreateBookingRequest struct {
	CustomerId int
	FlightId   int
	Code       string
	Status     string
	BookedDate int
}

type FindBookingRequest struct {
	Id uuid.UUID
}

type ListBookingRequest struct {
	CustomerId string
}
