package handlers

import (
	"booking/grpc/booking-grpc/models"
	"booking/grpc/booking-grpc/repositories"
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

type BookingHandler struct {
	pb.UnimplementedFPTBookingServer
	bookingRepository repositories.BookingRepository
	mu                *sync.Mutex
}

func NewBookingHandler(bookingRepository repositories.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepository: bookingRepository,
		mu:                &sync.Mutex{},
	}, nil
}

func (h *BookingHandler) CreateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	pRequest := &models.Booking{
		Id:         uuid.New(),
		CustomerId: in.CustomerId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		Status:     "created",
		BookedDate: in.BookedDate.AsTime(),
	}

	booking, err := h.bookingRepository.CreateBooking(ctx, pRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pResponse := &pb.Booking{
		Id:         booking.Id.String(),
		CustomerId: booking.CustomerId,
		FlightId:   booking.FlightId,
		Code:       booking.Code,
		Status:     booking.Status,
		BookedDate: timestamppb.New(booking.BookedDate),
	}

	return pResponse, nil
}

func (h *BookingHandler) FindBooking(ctx context.Context, req *pb.FindBookingRequest) (*pb.Booking, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	res, err := h.bookingRepository.FindBooking(ctx, uuid.MustParse(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "booking is not found")
		}
		return nil, err
	}

	pRes := &pb.Booking{}
	err = copier.Copy(&pRes, &res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *BookingHandler) CancelBooking(ctx context.Context, in *pb.CancelBookingRequest) (*pb.Booking, error) {
	booking, err := h.bookingRepository.GetBookingByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	booking.Status = "Cancel"

	res, err := h.bookingRepository.UpdateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Booking{}
	err = copier.Copy(&pRes, &res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
