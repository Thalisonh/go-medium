package health

import (
	"github.com/gin-gonic/gin"
)

func Router(allGroup *gin.RouterGroup) {
	r := NewHealthController()

	allGroup.GET("", r.Health)
}
