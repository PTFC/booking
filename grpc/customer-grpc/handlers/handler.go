package handlers

import (
	"booking/grpc/customer-grpc/models"
	"booking/grpc/customer-grpc/repositories"
	"booking/grpc/customer-grpc/requests"
	"booking/pb"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"sync"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	pb.UnimplementedFPTCustomerServer
	customerRepository repositories.CustomerRepository
	mu                 *sync.Mutex
}

func NewCustomerHandler(customerRepository repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepository,
		mu:                 &sync.Mutex{},
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	pRequest := &models.Customer{
		Id:          uuid.New(),
		Name:        in.Name,
		LicenseId:   in.LicenseId,
		PhoneNumber: in.PhoneNumber,
		Email:       in.Email,
		Password:    in.Password,
		Active:      in.Active,
	}

	customer, err := h.customerRepository.CreateCustomer(ctx, pRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pResponse := &pb.Customer{
		Id:          customer.Id.String(),
		Name:        customer.Name,
		LicenseId:   customer.LicenseId,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Password:    customer.Password,
		Active:      customer.Active,
	}

	return pResponse, nil
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Name != "" {
		customer.Name = in.Name
	}

	if in.LicenseId != "" {
		customer.LicenseId = in.LicenseId
	}

	if in.PhoneNumber != "" {
		customer.PhoneNumber = in.PhoneNumber
	}

	if in.Email != "" {
		customer.Email = in.Email
	}

	if in.Password != "" {
		secretKey := viper.GetString("postgres.username")
		h := hmac.New(sha256.New, []byte(secretKey))
		h.Write([]byte(in.Password))
		sha := hex.EncodeToString(h.Sum(nil))
		customer.Password = sha
	}

	if in.Active != customer.Active {
		customer.Active = in.Active
	}

	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.Customer{
		Id:          newCustomer.Id.String(),
		Name:        newCustomer.Name,
		LicenseId:   newCustomer.LicenseId,
		PhoneNumber: newCustomer.PhoneNumber,
		Email:       newCustomer.Email,
		Password:    newCustomer.Password,
		Active:      newCustomer.Active,
	}

	return pResponse, nil
}

func (h *CustomerHandler) BookingHistory(ctx context.Context, in *pb.BookingHistoryRequest) (*pb.BookingHistoryResponse, error) {
	bookings, err := h.customerRepository.BookingHistory(ctx, &requests.BookingHistoryRequest{
		CustomerId: in.CustomerId,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.BookingHistoryResponse{
		Bookings: []*pb.Booking{},
	}

	err = copier.CopyWithOption(&pRes.Bookings, &bookings, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})

	if err != nil {
		return nil, err
	}

	return pRes, nil
}
