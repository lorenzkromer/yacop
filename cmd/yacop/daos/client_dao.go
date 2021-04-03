package daos

import (
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/fitchlol/yacop/cmd/yacop/models"
)

// VehicleDAO persists client data in database
type VehicleDAO struct{}

// NewVehicleDAO creates a new VehicleDAO
func NewVehicleDAO() *VehicleDAO {
	return &VehicleDAO{}
}

func (dao *VehicleDAO) Create(client models.Vehicle) (*models.Vehicle, error) {
	err := config.Config.DB.Save(&client).Error

	return &client, err
}

func (dao *VehicleDAO) GetAll() ([]*models.Vehicle, error) {
	var clients []*models.Vehicle

	err := config.Config.DB.Find(&clients).
		Error

	return clients, err
}

func (dao *VehicleDAO) GetById(id string) (*models.Vehicle, error) {
	var client models.Vehicle

	err := config.Config.DB.Where("id = ?", id).
		First(&client).
		Error

	return &client, err
}

func (dao *VehicleDAO) Update(client models.Vehicle) (*models.Vehicle, error) {
	err := config.Config.DB.Model(&client).
		Updates(&client).
		Error

	return &client, err
}

func (dao *VehicleDAO) Delete(client models.Vehicle) error {
	err := config.Config.DB.Delete(&client).Error

	return err
}
