package routes

import (
	"github.com/fitchlol/yacop/cmd/yacop/apis"
	"github.com/gin-gonic/gin"
)

func DummyRoutesRegister(router *gin.RouterGroup) {
	router.GET("/hello-world", apis.SayHelloWorld)
}
