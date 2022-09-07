package requests

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	LicenseId   string `json:"license_id" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Active      bool   `json:"active" binding:"required"`
}

type UpdateCustomerRequest struct {
	Id          string `json:"id" binding:"required"`
	Name        string `json:"name"`
	LicenseId   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Active      bool   `json:"active"`
}

type ChangePasswordRequest struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BookingHistoryRequest struct {
	CustomerId int64 `json:"id" binding:"required"`
}
