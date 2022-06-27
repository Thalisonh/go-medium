package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthController interface {
	Health(ctx *gin.Context)
}

type HealthController struct {
}

func NewHealthController() IHealthController {
	return &HealthController{}
}

func (c *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
