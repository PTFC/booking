package handlers

import (
	"booking/api/customer-api/requests"
	"booking/api/customer-api/responses"
	"booking/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
	BookingHistory(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.FPTCustomerClient
}

func NewCustomerHandler(customerClient pb.FPTCustomerClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	req := requests.CreateCustomerRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Customer{
		Name:        req.Name,
		LicenseId:   req.LicenseId,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		Active:      req.Active,
	}

	pRes, err := h.customerClient.CreateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		ID:          pRes.Id,
		Name:        pRes.Name,
		LicenseId:   pRes.LicenseId,
		PhoneNumber: pRes.PhoneNumber,
		Email:       pRes.Email,
		Active:      pRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	req := requests.UpdateCustomerRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Customer{
		Id:          req.Id,
		Name:        req.Name,
		LicenseId:   req.LicenseId,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Active:      req.Active,
	}

	pRes, err := h.customerClient.UpdateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		ID:          pRes.Id,
		Name:        pRes.Name,
		LicenseId:   pRes.LicenseId,
		PhoneNumber: pRes.PhoneNumber,
		Email:       pRes.Email,
		Active:      pRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) ChangePassword(c *gin.Context) {
	req := requests.ChangePasswordRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Customer{
		Id:       req.Id,
		Password: req.Password,
	}

	pRes, err := h.customerClient.UpdateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		ID:          pRes.Id,
		Name:        pRes.Name,
		LicenseId:   pRes.LicenseId,
		PhoneNumber: pRes.PhoneNumber,
		Email:       pRes.Email,
		Active:      pRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) BookingHistory(c *gin.Context) {
	req := requests.BookingHistoryRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.BookingHistoryRequest{
		CustomerId: req.CustomerId,
	}

	pRes, err := h.customerClient.BookingHistory(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := []*responses.BookingResponse{}
	for _, booking := range pRes.Bookings {
		dto = append(dto, &responses.BookingResponse{
			ID:         booking.Id,
			CustomerId: booking.CustomerId,
			FlightId:   booking.FlightId,
			Code:       booking.Code,
			Status:     booking.Status,
			BookedDate: booking.BookedDate.AsTime(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
