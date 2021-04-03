package daos

import (
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"gorm.io/gorm/clause"
)

// VehicleDAO persists vehicle data in database
type VehicleDAO struct{}

// NewVehicleDAO creates a new VehicleDAO
func NewVehicleDAO() *VehicleDAO {
	return &VehicleDAO{}
}

func (dao *VehicleDAO) Create(vehicle models.Vehicle) (*models.Vehicle, error) {
	err := config.Config.DB.Save(&vehicle).Error

	return &vehicle, err
}

func (dao *VehicleDAO) GetAll() ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle

	err := config.Config.DB.Preload(clause.Associations).
		Find(&vehicles).
		Error

	return vehicles, err
}

func (dao *VehicleDAO) GetById(id string) (*models.Vehicle, error) {
	var vehicle models.Vehicle

	err := config.Config.DB.Where("id = ?", id).
		First(&vehicle).
		Error

	return &vehicle, err
}

func (dao *VehicleDAO) Update(vehicle models.Vehicle) (*models.Vehicle, error) {
	err := config.Config.DB.Model(&vehicle).
		Updates(&vehicle).
		Error

	return &vehicle, err
}

func (dao *VehicleDAO) Delete(vehicle models.Vehicle) error {
	err := config.Config.DB.Delete(&vehicle).Error

	return err
}
