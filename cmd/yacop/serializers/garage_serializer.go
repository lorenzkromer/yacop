package serializers

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type GarageSerializer struct {
	C           *gin.Context
	GarageModel *models.Garage
}

type GarageResponse struct {
	ID        string            `json:"id"`
	Vehicles  []VehicleResponse `json:"vehicles"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

func (s *GarageSerializer) Response() GarageResponse {
	var vehicleResponses []VehicleResponse
	for _, vehicle := range s.GarageModel.Vehicles {
		vehicleSerializer := VehicleSerializer{C: s.C, VehicleModel: &vehicle}
		vehicleResponses = append(vehicleResponses, vehicleSerializer.Response())
	}
	return GarageResponse{
		ID:        s.GarageModel.ID,
		Vehicles:  vehicleResponses,
		CreatedAt: s.GarageModel.CreatedAt,
		UpdatedAt: s.GarageModel.UpdatedAt,
		DeletedAt: s.GarageModel.DeletedAt,
	}
}
