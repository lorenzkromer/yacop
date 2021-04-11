package validators

import (
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
)

type ManufacturerModelValidator struct {
	Manufacturer struct {
		Name string `json:"name" binding:"required,min=1,max=255"`
	}
	ManufacturerModel models.Manufacturer `json:"-"`
}

func (s *ManufacturerModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.ManufacturerModel.Name = s.Manufacturer.Name
	return nil
}

// You can put the default value of a Validator here
func NewManufacturerModelValidator() ManufacturerModelValidator {
	validator := ManufacturerModelValidator{}
	return validator
}

func NewManufacturerModelValidatorFillWith(manufacturer models.Manufacturer) ManufacturerModelValidator {
	validator := NewManufacturerModelValidator()
	validator.Manufacturer.Name = manufacturer.Name
	return validator
}
