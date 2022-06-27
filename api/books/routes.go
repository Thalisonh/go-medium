package books

import (
	"github.com/Thalisonh/go-medium/database"
	"github.com/gin-gonic/gin"
)

func Router(allGroup *gin.RouterGroup) {
	r := NewBookRepository(database.GetDb())
	c := NewBookController(r)

	allGroup.POST("", c.CreateBook)
	allGroup.GET("", c.ShowBooks)
}
