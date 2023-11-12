package models

import (
	"github.com/KarkiAnmol/BookStore-API-Golang/pkg/config"
	"github.com/jinzhu/gorm"
)

// db holds the reference to the GORM database instance.
var db *gorm.DB

// Book represents the model for a book in the BookStore.
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init initializes the database connection and performs auto-migration for the Book model.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook creates a new book record in the database.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetBooks retrieves all books from the database.
func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById retrieves a book by its ID from the database.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}
