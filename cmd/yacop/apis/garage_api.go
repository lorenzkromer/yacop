package apis

import (
	"github.com/fitchlol/yacop/cmd/yacop/daos"
	middlewares "github.com/fitchlol/yacop/cmd/yacop/middleware"
	"github.com/fitchlol/yacop/cmd/yacop/serializers"
	"github.com/fitchlol/yacop/cmd/yacop/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GarageInit(c *gin.Context) {
	s := services.NewGarageService(daos.NewGarageDAO())
	user := middlewares.GetUserContext(c)
	if garage, err := s.Register(user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err)
		log.Error(err)
	} else {
		serializer := serializers.GarageSerializer{C: c, GarageModel: garage}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

func GarageByUser(c *gin.Context) {
	s := services.NewGarageService(daos.NewGarageDAO())
	user := middlewares.GetUserContext(c)
	if garage, err := s.GetByUserId(user.ID); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error(err)
	} else {
		serializer := serializers.GarageSerializer{C: c, GarageModel: garage}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
