package routes

import (
	"github.com/fitchlol/yacop/cmd/yacop/apis"
	"github.com/gin-gonic/gin"
)

func ManufacturersRoutesRegister(router *gin.RouterGroup) {
	router.POST("/", apis.ManufacturerCreate)
	router.GET("/", apis.Manufacturers)
	router.GET("/:id", apis.ManufacturerById)
	router.PUT("/:id", apis.ManufacturerUpdate)
	router.DELETE("/:id", apis.ManufacturerDelete)
}
