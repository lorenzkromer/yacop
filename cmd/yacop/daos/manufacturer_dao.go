package daos

import (
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/fitchlol/yacop/cmd/yacop/models"
)

// ManufacturerDAO persists manufacturer data in database
type ManufacturerDAO struct{}

// NewManufacturerDAO creates a new ManufacturerDAO
func NewManufacturerDAO() *ManufacturerDAO {
	return &ManufacturerDAO{}
}

func (dao *ManufacturerDAO) Create(manufacturer models.Manufacturer) (*models.Manufacturer, error) {
	err := config.Config.DB.Save(&manufacturer).Error

	return &manufacturer, err
}

func (dao *ManufacturerDAO) GetAll() ([]*models.Manufacturer, error) {
	var manufacturers []*models.Manufacturer

	err := config.Config.DB.Find(&manufacturers).
		Error

	return manufacturers, err
}

func (dao *ManufacturerDAO) GetById(id string) (*models.Manufacturer, error) {
	var m models.Manufacturer

	err := config.Config.DB.Where("id = ?", id).
		First(&m).
		Error

	return &m, err
}

func (dao *ManufacturerDAO) Update(manufacturer models.Manufacturer) (*models.Manufacturer, error) {
	err := config.Config.DB.Model(&manufacturer).
		Updates(&manufacturer).
		Error

	return &manufacturer, err
}

func (dao *ManufacturerDAO) Delete(manufacturer models.Manufacturer) error {
	err := config.Config.DB.Delete(&manufacturer).Error

	return err
}
