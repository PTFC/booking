package repositories

import (
	"booking/database"
	"booking/grpc/flight-grpc/models"
	"booking/grpc/flight-grpc/requests"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type FlightRepository interface {
	GetFlightByID(context context.Context, id uuid.UUID) (*models.Flight, error)
	CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	SearchFlights(ctx context.Context, req *requests.SearchFlightRequest) ([]*models.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) GetFlightByID(ctx context.Context, id uuid.UUID) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Where(&models.Flight{Id: id}).First(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Where(&models.Flight{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) SearchFlights(ctx context.Context, req *requests.SearchFlightRequest) ([]*models.Flight, error) {
	fromFlights := []*models.Flight{}
	if err := m.Where(&models.Flight{From: req.QueryString}).Find(&fromFlights).Error; err != nil {
		return nil, err
	}

	toFlights := []*models.Flight{}
	if err := m.Where(&models.Flight{To: req.QueryString}).Find(&toFlights).Error; err != nil {
		return nil, err
	}

	dateFlights := []*models.Flight{}
	if err := m.Where(&models.Flight{Date: req.QueryString}).Find(&dateFlights).Error; err != nil {
		return nil, err
	}

	flights := append(fromFlights, toFlights...)
	flights = append(flights, dateFlights...)

	return flights, nil
}
