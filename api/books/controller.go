package books

import (
	"net/http"

	"github.com/Thalisonh/go-medium/model"
	"github.com/gin-gonic/gin"
)

type ControllerBook struct {
	repository IBookRepository
}

func NewBookController(repository IBookRepository) ControllerBook {
	return ControllerBook{repository: repository}
}

func (s *ControllerBook) CreateBook(c *gin.Context) {
	var newBook model.Book

	errJson := c.ShouldBindJSON(&newBook)
	if errJson != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": errJson,
		})
	}

	book, err := s.repository.CreateBook(&newBook)

	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
