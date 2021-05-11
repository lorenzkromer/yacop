package validators

import (
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
)

type GarageModelValidator struct {
	Garage struct {
		UserID string `json:"user_id" binding:"required,uuid"`
	}
	GarageModel models.Garage `json:"-"`
}

func (s *GarageModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.GarageModel.UserID = s.Garage.UserID
	return nil
}

// You can put the default value of a Validator here
func NewGarageModelValidator() GarageModelValidator {
	validator := GarageModelValidator{}
	return validator
}

func NewGarageModelValidatorFillWith(manufacturer models.Garage) GarageModelValidator {
	validator := NewGarageModelValidator()
	validator.Garage.UserID = manufacturer.UserID
	return validator
}
