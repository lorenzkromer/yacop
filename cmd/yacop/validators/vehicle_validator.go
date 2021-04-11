package validators

import (
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/enums"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
)

type VehicleModelValidator struct {
	Vehicle struct {
		FullName                 string         `json:"full_name" binding:"required,min=4,max=255"`
		FuelType                 enums.FuelType `json:"fuel_type" binding:"required,min=4,max=255"`
		MaximumKilometersPerHour int            `json:"maximum_fuel_type" binding:"required,gt=0"`
		MaximumKilowatts         int            `json:"maximum_kilowatts" binding:"required,gt=0"`
		WeightInKilograms        int            `json:"weight_in_kilograms" binding:"required,gt=0"`
		ManufacturerID           string         `json:"manufacturer_id" binding:"required,uuid"`
	}
	VehicleModel models.Vehicle `json:"-"`
}

func (s *VehicleModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.VehicleModel.FullName = s.Vehicle.FullName
	s.VehicleModel.FuelType = s.Vehicle.FuelType
	s.VehicleModel.MaximumKilowatts = s.Vehicle.MaximumKilowatts
	s.VehicleModel.MaximumKilometersPerHour = s.Vehicle.MaximumKilometersPerHour
	s.VehicleModel.WeightInKilograms = s.Vehicle.WeightInKilograms
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
	validator.Vehicle.FuelType = model.FuelType
	validator.Vehicle.MaximumKilowatts = model.MaximumKilowatts
	validator.Vehicle.MaximumKilometersPerHour = model.MaximumKilometersPerHour
	validator.Vehicle.WeightInKilograms = model.WeightInKilograms
	validator.Vehicle.ManufacturerID = model.ManufacturerID
	return validator
}
