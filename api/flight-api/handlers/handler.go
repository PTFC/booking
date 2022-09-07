package handlers

import (
	"booking/api/flight-api/requests"
	"booking/api/flight-api/responses"
	"booking/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	SearchFlight(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FPTFlightClient
}

func NewFlightHandler(flightClient pb.FPTFlightClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := requests.CreateFlightRequest{}

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

	pReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(req.Date),
		AvailableSlot: req.AvailableSlot,
	}

	pRes, err := h.flightClient.CreateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		Date:          pRes.Date.AsTime(),
		Status:        pRes.Status,
		AvailableSlot: pRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) UpdateFlight(c *gin.Context) {
	req := requests.UpdateFlightRequest{}

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

	pReq := &pb.Flight{
		Id:            req.Id,
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(req.Date),
		AvailableSlot: req.AvailableSlot,
	}

	pRes, err := h.flightClient.UpdateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		Date:          pRes.Date.AsTime(),
		Status:        pRes.Status,
		AvailableSlot: pRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) SearchFlight(c *gin.Context) {
	req := requests.SearchFlightRequest{}

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

	pReq := &pb.SearchFlightRequest{
		QueryString: req.QueryString,
	}

	pRes, err := h.flightClient.SearchFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := []*responses.FlightResponse{}
	for _, flight := range pRes.Flights {
		dto = append(dto, &responses.FlightResponse{
			ID:            flight.Id,
			Name:          flight.Name,
			From:          flight.From,
			To:            flight.To,
			Date:          flight.Date.AsTime(),
			Status:        flight.Status,
			AvailableSlot: flight.AvailableSlot,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
