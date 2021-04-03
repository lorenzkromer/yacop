package routes

import (
	"github.com/fitchlol/yacop/cmd/yacop/apis"
	"github.com/gin-gonic/gin"
)

func VehiclesRoutesRegister(router *gin.RouterGroup) {
	router.POST("/", apis.VehicleCreate)
	router.GET("/", apis.Vehicles)
	router.GET("/:id", apis.VehicleById)
	router.PUT("/:id", apis.VehicleUpdate)
	router.DELETE("/:id", apis.VehicleDelete)
}
