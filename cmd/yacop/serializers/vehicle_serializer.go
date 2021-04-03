package serializers

import (
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
	ID        string         `json:"id"`
	FullName  string         `json:"full_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (s *VehicleSerializer) Response() VehicleResponse {
	return VehicleResponse{
		ID:        s.VehicleModel.ID,
		FullName:  s.VehicleModel.FullName,
		CreatedAt: s.VehicleModel.CreatedAt,
		UpdatedAt: s.VehicleModel.UpdatedAt,
		DeletedAt: s.VehicleModel.DeletedAt,
	}
}
