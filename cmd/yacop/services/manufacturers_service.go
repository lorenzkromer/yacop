package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
)

type manufacturerDAO interface {
	Create(manufacturer models.Manufacturer) (*models.Manufacturer, error)
	GetAll() ([]*models.Manufacturer, error)
	GetById(id string) (*models.Manufacturer, error)
	Update(manufacturer models.Manufacturer) (*models.Manufacturer, error)
	Delete(manufacturer models.Manufacturer) error
}

type ManufacturersService struct {
	dManufacturer manufacturerDAO
}

func NewManufacturersService(dManufacturer manufacturerDAO) *ManufacturersService {
	return &ManufacturersService{dManufacturer}
}

func (s *ManufacturersService) Create(manufacturer models.Manufacturer) (*models.Manufacturer, error) {
	return s.dManufacturer.Create(manufacturer)
}

func (s *ManufacturersService) GetAll() ([]*models.Manufacturer, error) {
	return s.dManufacturer.GetAll()
}

func (s *ManufacturersService) GetById(id string) (*models.Manufacturer, error) {
	return s.dManufacturer.GetById(id)
}

func (s *ManufacturersService) Update(manufacturer models.Manufacturer) (*models.Manufacturer, error) {
	return s.dManufacturer.Update(manufacturer)
}

func (s *ManufacturersService) Delete(manufacturer models.Manufacturer) error {
	return s.dManufacturer.Delete(manufacturer)
}
