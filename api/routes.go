package api

import (
	"github.com/Thalisonh/go-medium/api/books"
	"github.com/Thalisonh/go-medium/api/health"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	req := r.Group("/api")

	health.Router(req.Group("/health"))
	books.Router(req.Group("/books"))
}
