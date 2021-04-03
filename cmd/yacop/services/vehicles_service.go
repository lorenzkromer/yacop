package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
)

type vehicleDAO interface {
	Create(vehicle models.Vehicle) (*models.Vehicle, error)
	GetAll() ([]*models.Vehicle, error)
	GetById(id string) (*models.Vehicle, error)
	Update(vehicle models.Vehicle) (*models.Vehicle, error)
	Delete(vehicle models.Vehicle) error
}

type VehiclesService struct {
	dVehicle vehicleDAO
}

func NewVehiclesService(dVehicle vehicleDAO) *VehiclesService {
	return &VehiclesService{dVehicle}
}

func (s *VehiclesService) Create(vehicle models.Vehicle) (*models.Vehicle, error) {
	return s.dVehicle.Create(vehicle)
}

func (s *VehiclesService) GetAll() ([]*models.Vehicle, error) {
	return s.dVehicle.GetAll()
}

func (s *VehiclesService) GetById(id string) (*models.Vehicle, error) {
	return s.dVehicle.GetById(id)
}

func (s *VehiclesService) Update(vehicle models.Vehicle) (*models.Vehicle, error) {
	return s.dVehicle.Update(vehicle)
}

func (s *VehiclesService) Delete(vehicle models.Vehicle) error {
	return s.dVehicle.Delete(vehicle)
}
