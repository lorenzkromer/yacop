package serializers

import (
	"github.com/fitchlol/yacop/cmd/yacop/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ManufacturerSerializer struct {
	C         *gin.Context
	ManufacturerModel *models.Manufacturer
}

type ManufacturerResponse struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (s *ManufacturerSerializer) Response() ManufacturerResponse {
	return ManufacturerResponse{
		ID:        s.ManufacturerModel.ID,
		Name:      s.ManufacturerModel.Name,
		CreatedAt: s.ManufacturerModel.CreatedAt,
		UpdatedAt: s.ManufacturerModel.UpdatedAt,
		DeletedAt: s.ManufacturerModel.DeletedAt,
	}
}
