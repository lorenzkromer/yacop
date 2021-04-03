package serializers

import (
	"github.com/gin-gonic/gin"
)

type DummySerializer struct {
	C *gin.Context
}

type DummyResponse struct {
	Message string `json:"message"`
}

func (s *DummySerializer) Response() DummyResponse {
	return DummyResponse{
		Message: "Hello World",
	}
}
