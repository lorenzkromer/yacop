package routes

import (
	"github.com/fitchlol/yacop/cmd/yacop/apis"
	"github.com/gin-gonic/gin"
)

func GarageRoutesRegister(router *gin.RouterGroup) {
	router.GET("", apis.GarageByUser)
	// router.GET("/init", apis.GarageInit)

	router.POST("/vehicles", apis.VehicleCreate)
	router.GET("/vehicles", apis.Vehicles)
	router.GET("/vehicles/:id", apis.VehicleById)
	router.PUT("/vehicles/:id", apis.VehicleUpdate)
	router.DELETE("/vehicles/:id", apis.VehicleDelete)
}
