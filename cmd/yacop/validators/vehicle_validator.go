package validators

import (
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
)

type VehicleModelValidator struct {
	Vehicle struct {
		FullName       string `json:"full_name" binding:"required,min=4,max=255"`
		ManufacturerID string `json:"manufacturer_id" binding:"required,uuid"`
	}
	VehicleModel models.Vehicle `json:"-"`
}

func (s *VehicleModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.VehicleModel.FullName = s.Vehicle.FullName
	s.VehicleModel.ManufacturerID = s.Vehicle.ManufacturerID
	return nil
}

// You can put the default value of a Validator here
func NewVehicleModelValidator() VehicleModelValidator {
	validator := VehicleModelValidator{}
	return validator
}

func NewVehicleModelValidatorFillWith(model models.Vehicle) VehicleModelValidator {
	validator := NewVehicleModelValidator()
	validator.Vehicle.FullName = model.FullName
	validator.Vehicle.ManufacturerID = model.ManufacturerID
	return validator
}
