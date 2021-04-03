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

func ManufacturerCreate(c *gin.Context) {
	v := validators.NewManufacturerModelValidator()
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	s := services.NewManufacturersService(daos.NewManufacturerDAO())
	if m, err := s.Create(v.ManufacturerModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
		log.Error(err)
	} else {
		serializer := serializers.ManufacturerSerializer{C: c, ManufacturerModel: m}
		c.JSON(http.StatusCreated, serializer.Response())
	}
}

func Manufacturers(c *gin.Context) {
	s := services.NewManufacturersService(daos.NewManufacturerDAO())
	if users, err := s.GetAll(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Error(err)
	} else {
		var makesResponse []serializers.ManufacturerResponse
		for _, m := range users {
			serializer := serializers.ManufacturerSerializer{C: c, ManufacturerModel: m}
			makesResponse = append(makesResponse, serializer.Response())
		}
		c.JSON(http.StatusOK, makesResponse)
	}
}

func ManufacturerById(c *gin.Context) {
	s := services.NewManufacturersService(daos.NewManufacturerDAO())
	id := c.Param("id")
	if m, err := s.GetById(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error(err)
	} else {
		serializer := serializers.ManufacturerSerializer{C: c, ManufacturerModel: m}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

func ManufacturerUpdate(c *gin.Context) {
	s := services.NewManufacturersService(daos.NewManufacturerDAO())
	id := c.Param("id")
	if databaseManufacturers, err := s.GetById(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
		log.Error(err)
	} else {
		v := validators.NewManufacturerModelValidatorFillWith(*databaseManufacturers)
		if err := v.Bind(c); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			log.Error(err)
			return
		}
		v.ManufacturerModel.ID = databaseManufacturers.ID
		if updatedManufacturers, err := s.Update(v.ManufacturerModel); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			log.Error(err)
			return
		} else {
			serializer := serializers.ManufacturerSerializer{C: c, ManufacturerModel: updatedManufacturers}
			c.JSON(http.StatusOK, serializer.Response())
		}
	}
}

func ManufacturerDelete(c *gin.Context) {
	s := services.NewManufacturersService(daos.NewManufacturerDAO())
	id := c.Param("id")
	if databaseManufacturers, err := s.GetById(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error(err)
	} else {
		if err := s.Delete(*databaseManufacturers); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Error(err)
		} else {
			c.JSON(http.StatusNoContent, nil)
		}
	}
}

