package requests

import "github.com/google/uuid"

type CreateFlightRequest struct {
	Name          string
	From          string
	To            string
	Date          int
	AvailableSlot int
}

type UpdateFlightRequest struct {
	Id            uuid.UUID
	Name          string
	From          string
	To            string
	Date          int
	AvailableSlot int
}

type SearchFlightRequest struct {
	QueryString string
}
