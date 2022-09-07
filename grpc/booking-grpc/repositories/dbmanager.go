package repositories

import (
	"booking/database"
	"booking/grpc/booking-grpc/models"
	"booking/grpc/booking-grpc/requests"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type BookingRepository interface {
	GetBookingByID(context context.Context, id uuid.UUID) (*models.Booking, error)
	CreateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
	FindBooking(ctx context.Context, id uuid.UUID) (*models.Booking, error)
	ListBooking(ctx context.Context, req *requests.ListBookingRequest) ([]*models.Booking, error)
	UpdateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Booking{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) GetBookingByID(ctx context.Context, id uuid.UUID) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.Where(&models.Booking{Id: id}).First(&booking).Error; err != nil {
		return nil, err
	}

	return &booking, nil
}

func (m *dbmanager) CreateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) FindBooking(ctx context.Context, id uuid.UUID) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.Where(&models.Booking{Id: id}).First(&booking).Error; err != nil {
		return nil, err
	}

	return &booking, nil
}

func (m *dbmanager) ListBooking(ctx context.Context, req *requests.ListBookingRequest) ([]*models.Booking, error) {
	bookings := []*models.Booking{}

	if err := m.Where("customer_id = ?", req.CustomerId).Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *dbmanager) UpdateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	if err := m.Where(&models.Booking{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
