package daos

import (
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/fitchlol/yacop/cmd/yacop/models"
)

// GarageDAO persists Garage data in database
type GarageDAO struct{}

// NewGarageDAO creates a new GarageDAO
func NewGarageDAO() *GarageDAO {
	return &GarageDAO{}
}

func (dao *GarageDAO) Create(garage models.Garage) (*models.Garage, error) {
	err := config.Config.DB.Save(&garage).Error

	return &garage, err
}

func (dao *GarageDAO) GetByUserId(userId string) (*models.Garage, error) {
	var garage models.Garage

	err := config.Config.DB.Where("user_id = ?", userId).
		First(&garage).
		Error

	return &garage, err
}

func (dao *GarageDAO) Update(garage models.Garage) (*models.Garage, error) {
	err := config.Config.DB.Model(&garage).
		Updates(&garage).
		Error

	return &garage, err
}

func (dao *GarageDAO) Delete(garage models.Garage) error {
	err := config.Config.DB.Delete(&garage).Error

	return err
}
