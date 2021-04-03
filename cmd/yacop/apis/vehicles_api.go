package apis

import (
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/daos"
	"github.com/fitchlol/yacop/cmd/yacop/serializers"
	"github.com/fitchlol/yacop/cmd/yacop/services"
	"github.com/fitchlol/yacop/cmd/yacop/validators"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func VehicleCreate(c *gin.Context) {
	v := validators.NewVehicleModelValidator()
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	s := services.NewVehiclesService(daos.NewVehicleDAO())
	if client, err := s.Create(v.VehicleModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
		log.Error(err)
	} else {
		serializer := serializers.VehicleSerializer{C: c, VehicleModel: client}
		c.JSON(http.StatusCreated, serializer.Response())
	}
}

func Vehicles(c *gin.Context) {
	s := services.NewVehiclesService(daos.NewVehicleDAO())
	if users, err := s.GetAll(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Error(err)
	} else {
		var vehiclesResponse []serializers.VehicleResponse
		for _, vehicle := range users {
			serializer := serializers.VehicleSerializer{C: c, VehicleModel: vehicle}
			vehiclesResponse = append(vehiclesResponse, serializer.Response())
		}
		c.JSON(http.StatusOK, vehiclesResponse)
	}
}

func VehicleById(c *gin.Context) {
	s := services.NewVehiclesService(daos.NewVehicleDAO())
	id := c.Param("id")
	if client, err := s.GetById(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error(err)
	} else {
		serializer := serializers.VehicleSerializer{C: c, VehicleModel: client}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

func VehicleUpdate(c *gin.Context) {
	s := services.NewVehiclesService(daos.NewVehicleDAO())
	id := c.Param("id")
	if databaseVehicle, err := s.GetById(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
		log.Error(err)
	} else {
		v := validators.NewVehicleModelValidatorFillWith(*databaseVehicle)
		if err := v.Bind(c); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			log.Error(err)
			return
		}
		v.VehicleModel.ID = databaseVehicle.ID
		if updatedVehicle, err := s.Update(v.VehicleModel); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			log.Error(err)
			return
		} else {
			serializer := serializers.VehicleSerializer{C: c, VehicleModel: updatedVehicle}
			c.JSON(http.StatusOK, serializer.Response())
		}
	}
}

func VehicleDelete(c *gin.Context) {
	s := services.NewVehiclesService(daos.NewVehicleDAO())
	id := c.Param("id")
	if databaseVehicle, err := s.GetById(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error(err)
	} else {
		if err := s.Delete(*databaseVehicle); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Error(err)
		} else {
			c.JSON(http.StatusNoContent, nil)
		}
	}
}

