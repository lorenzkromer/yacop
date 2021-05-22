package services

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"gorm.io/gorm"
)

type vehicleDAO interface {
	Create(vehicle models.Vehicle) (*models.Vehicle, error)
	GetByGarage(userId string) ([]*models.Vehicle, error)
	GetByGarageAndId(garageId string, id string) (*models.Vehicle, error)
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

func (s *VehiclesService) GetByGarage(userId string) ([]*models.Vehicle, error) {
	return s.dVehicle.GetByGarage(userId)
}

func (s *VehiclesService) GetByGarageAndId(garage models.Garage, id string) (vehicle *models.Vehicle, err error) {
	for _, v := range garage.Vehicles {
		if v.ID == id {
			vehicle = &v
			break
		}
	}
	if vehicle == nil {
		return nil, gorm.ErrRecordNotFound
	}
	return vehicle, err
}

func (s *VehiclesService) Update(vehicle models.Vehicle) (*models.Vehicle, error) {
	return s.dVehicle.Update(vehicle)
}

func (s *VehiclesService) Delete(vehicle models.Vehicle) error {
	return s.dVehicle.Delete(vehicle)
}
