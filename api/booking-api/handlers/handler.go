package handlers

import (
	"booking/api/booking-api/requests"
	"booking/api/booking-api/responses"
	"booking/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler interface {
	CreateBooking(c *gin.Context)
	FindBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.FPTBookingClient
}

func NewBookingHandler(bookingClient pb.FPTBookingClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}

func (h *bookingHandler) CreateBooking(c *gin.Context) {
	req := requests.CreateBookingRequest{}

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

	pReq := &pb.Booking{
		CustomerId: req.CustomerId,
		FlightId:   req.FlightId,
		Code:       req.Code,
		Status:     req.Status,
		BookedDate: timestamppb.New(req.BookedDate),
	}

	pRes, err := h.bookingClient.CreateBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{
		ID:         pRes.Id,
		CustomerId: pRes.CustomerId,
		FlightId:   pRes.FlightId,
		Code:       pRes.Code,
		Status:     pRes.Status,
		BookedDate: pRes.BookedDate.AsTime(),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *bookingHandler) FindBooking(c *gin.Context) {
	req := requests.FindBookingRequest{}

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

	pReq := &pb.FindBookingRequest{
		Id: req.Id,
	}

	pRes, err := h.bookingClient.FindBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{
		ID:         pRes.Id,
		CustomerId: pRes.CustomerId,
		FlightId:   pRes.FlightId,
		Code:       pRes.Code,
		Status:     pRes.Status,
		BookedDate: pRes.BookedDate.AsTime(),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *bookingHandler) CancelBooking(c *gin.Context) {
	req := requests.CancelBookingRequest{}

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

	pReq := &pb.CancelBookingRequest{
		Id: req.Id,
	}

	pRes, err := h.bookingClient.CancelBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{
		ID:         pRes.Id,
		CustomerId: pRes.CustomerId,
		FlightId:   pRes.FlightId,
		Code:       pRes.Code,
		Status:     pRes.Status,
		BookedDate: pRes.BookedDate.AsTime(),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
