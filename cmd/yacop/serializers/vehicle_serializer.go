package serializers

import (
	"github.com/fitchlol/yacop/cmd/yacop/enums"
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type VehicleSerializer struct {
	C            *gin.Context
	VehicleModel *models.Vehicle
}

type VehicleResponse struct {
	ID                       string               `json:"id"`
	FullName                 string               `json:"full_name"`
	FuelType                 enums.FuelType       `json:"fuel_type"`
	MaximumKilometersPerHour int                  `json:"maximum_fuel_type"`
	MaximumKilowatts         int                  `json:"maximum_kilowatts"`
	WeightInKilograms        int                  `json:"weight_in_kilograms"`
	Manufacturer             ManufacturerResponse `json:"manufacturer"`
	CreatedAt                time.Time            `json:"created_at"`
	UpdatedAt                time.Time            `json:"updated_at"`
	DeletedAt                gorm.DeletedAt       `json:"deleted_at"`
}

func (s *VehicleSerializer) Response() VehicleResponse {
	manufacturerSerializer := ManufacturerSerializer{C: s.C, ManufacturerModel: &s.VehicleModel.Manufacturer}
	return VehicleResponse{
		ID:                       s.VehicleModel.ID,
		FullName:                 s.VehicleModel.FullName,
		FuelType:                 s.VehicleModel.FuelType,
		MaximumKilometersPerHour: s.VehicleModel.MaximumKilometersPerHour,
		MaximumKilowatts:         s.VehicleModel.MaximumKilowatts,
		WeightInKilograms:        s.VehicleModel.WeightInKilograms,
		Manufacturer:             manufacturerSerializer.Response(),
		CreatedAt:                s.VehicleModel.CreatedAt,
		UpdatedAt:                s.VehicleModel.UpdatedAt,
		DeletedAt:                s.VehicleModel.DeletedAt,
	}
}
