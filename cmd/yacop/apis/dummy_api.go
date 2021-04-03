package apis

import (
	"github.com/fitchlol/yacop/cmd/yacop/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SayHelloWorld(c *gin.Context) {
	serializer := serializers.DummySerializer{C: c}
	c.JSON(http.StatusOK, serializer.Response())
}
