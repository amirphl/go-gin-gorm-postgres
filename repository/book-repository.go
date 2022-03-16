package repository

import (
	"github.com/amirphl/go-gin-gorm-postgres/entity"
	"gorm.io/gorm"
)

// BookRepository ...
type BookRepository interface {
	Create(book entity.Book) entity.Book
	Update(book entity.Book) entity.Book
	Retrive(bookID uint64) interface{}
	List() []entity.Book
	Delete(book entity.Book)
	IsOwner(userID uint64, bookID uint64) bool
}

type bookRepository struct {
	connection *gorm.DB
}

func (b *bookRepository) Create(book entity.Book) entity.Book {
	// TODO How does `userID` field gets set? Preload?
	b.connection.Save(&book)
	b.connection.Preload("User").Find(&book)
	return book
}

func (b *bookRepository) Update(book entity.Book) entity.Book {
	b.connection.Save(&book)
	b.connection.Preload("User").Find(&book)
	return book
}

func (b *bookRepository) Retrive(bookID uint64) interface{} {
	var book entity.Book

	res := b.connection.Find(&book, bookID)

	if res.Error != nil || (book == entity.Book{}) {
		return nil
	}

	return book
}

func (b *bookRepository) List() []entity.Book {
	var books []entity.Book
	b.connection.Preload("User").Find(&books)
	return books
}

func (b *bookRepository) Delete(book entity.Book) {
	b.connection.Delete(&book)
}

func (b *bookRepository) IsOwner(userID uint64, bookID uint64) bool {
	if book, ok := b.Retrive(bookID).(entity.Book); ok {
		return book.UserID == userID
	}

	return false
}

// CreateBookRepo ...
func CreateBookRepo(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}
