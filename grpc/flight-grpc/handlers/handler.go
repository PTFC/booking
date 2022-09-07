package handlers

import (
	"booking/grpc/flight-grpc/models"
	"booking/grpc/flight-grpc/repositories"
	"booking/grpc/flight-grpc/requests"
	"booking/pb"
	"context"
	"database/sql"
	"sync"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	pb.UnimplementedFPTFlightServer
	flightRepository repositories.FlightRepository
	mu               *sync.Mutex
}

func NewFlightHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
		mu:               &sync.Mutex{},
	}, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	pRequest := &models.Flight{
		Id:            uuid.New(),
		Name:          in.Name,
		From:          in.From,
		To:            in.To,
		Date:          in.Date.AsTime(),
		Status:        "created",
		AvailableSlot: in.AvailableSlot,
	}

	flight, err := h.flightRepository.CreateFlight(ctx, pRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pResponse := &pb.Flight{
		Id:            flight.Id.String(),
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		Date:          timestamppb.New(flight.Date),
		Status:        flight.Status,
		AvailableSlot: flight.AvailableSlot,
	}

	return pResponse, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	flight, err := h.flightRepository.GetFlightByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Name != "" {
		flight.Name = in.Name
	}

	if in.From != "" {
		flight.From = in.From
	}

	if in.To != "" {
		flight.To = in.To
	}

	if in.Date != nil {
		flight.Date = in.Date.AsTime()
	}

	if in.Status != "" {
		flight.Status = in.Status
	}

	if in.AvailableSlot > 0 {
		flight.AvailableSlot = in.AvailableSlot
	}

	newFlight, err := h.flightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.Flight{
		Id:            newFlight.Id.String(),
		Name:          newFlight.Name,
		From:          newFlight.From,
		To:            newFlight.To,
		Date:          timestamppb.New(newFlight.Date),
		Status:        newFlight.Status,
		AvailableSlot: newFlight.AvailableSlot,
	}

	return pResponse, nil
}

func (h *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.ListFlightResponse, error) {
	flights, err := h.flightRepository.SearchFlights(ctx, &requests.SearchFlightRequest{
		QueryString: in.QueryString,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.ListFlightResponse{
		Flights: []*pb.Flight{},
	}

	err = copier.CopyWithOption(&pRes.Flights, &flights, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})

	if err != nil {
		return nil, err
	}

	return pRes, nil
}
