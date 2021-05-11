package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
)

type garageDao interface {
	Create(garage models.Garage) (*models.Garage, error)
	GetByUserId(userId string) (*models.Garage, error)
	Update(garage models.Garage) (*models.Garage, error)
	Delete(garage models.Garage) error
}

type GarageService struct {
	dGarage garageDao
}

func NewGarageService(dGarage garageDao) *GarageService {
	return &GarageService{dGarage}
}

func (s *GarageService) Register(userId string) (*models.Garage, error) {
	newGarage := models.Garage{UserID: userId}
	return s.dGarage.Create(newGarage)
}

func (s *GarageService) GetByUserId(userId string) (*models.Garage, error) {
	return s.dGarage.GetByUserId(userId)
}

func (s *GarageService) Update(garage models.Garage) (*models.Garage, error) {
	return s.dGarage.Update(garage)
}

func (s *GarageService) Delete(garage models.Garage) error {
	return s.dGarage.Delete(garage)
}
