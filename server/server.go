package server

import (
	"github.com/Thalisonh/go-medium/api"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()

	api.Router(r)

	return r
}
