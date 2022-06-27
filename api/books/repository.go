package books

import (
	"github.com/Thalisonh/go-medium/model"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book *model.Book) (*model.Book, error)
	GetBooks() (*[]model.Book, error)
}

type RepositoryBook struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &RepositoryBook{db: db}
}

func (r *RepositoryBook) CreateBook(book *model.Book) (*model.Book, error) {
	return book, r.db.Create(&book).Error
}

func (r *RepositoryBook) GetBooks() (*[]model.Book, error) {
	var books []model.Book
	return &books, r.db.Find(&books).Error
}
