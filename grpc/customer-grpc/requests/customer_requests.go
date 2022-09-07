package requests

import "github.com/google/uuid"

type CreateCustomerRequest struct {
	Name        string
	LicenseId   string
	PhoneNumber string
	Email       string
	Password    string
	Active      bool
}

type UpdateCustomerRequest struct {
	Id          uuid.UUID
	Name        string
	LicenseId   string
	PhoneNumber string
	Email       string
	Active      bool
}

type ChangePasswordRequest struct {
	Id       uuid.UUID
	Password string
}

type BookingHistoryRequest struct {
	CustomerId int64
}
