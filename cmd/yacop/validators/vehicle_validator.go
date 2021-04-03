package validators

import (
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
)

type VehicleModelValidator struct {
	Vehicle struct {
		FullName string `json:"full_name" binding:"required,min=4,max=255"`
	}
	VehicleModel models.Vehicle `json:"-"`
}

func (s *VehicleModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.VehicleModel.FullName = s.Vehicle.FullName
	return nil
}

// You can put the default value of a Validator here
func NewVehicleModelValidator() VehicleModelValidator {
	modelValidator := VehicleModelValidator{}
	return modelValidator
}

func NewVehicleModelValidatorFillWith(clientModel models.Vehicle) VehicleModelValidator {
	modelValidator := NewVehicleModelValidator()
	modelValidator.Vehicle.FullName = clientModel.FullName
	return modelValidator
}
