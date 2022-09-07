package repositories

import (
	"booking/database"
	"booking/grpc/customer-grpc/models"
	"booking/grpc/customer-grpc/requests"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type CustomerRepository interface {
	GetCustomerByID(context context.Context, id uuid.UUID) (*models.Customer, error)
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	BookingHistory(ctx context.Context, req *requests.BookingHistoryRequest) ([]*models.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) GetCustomerByID(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Id: id}).First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (m *dbmanager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Where(&models.Customer{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) BookingHistory(ctx context.Context, req *requests.BookingHistoryRequest) ([]*models.Booking, error) {
	bookings := []*models.Booking{}

	if err := m.Where(&models.Booking{CustomerId: req.CustomerId}).Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}
