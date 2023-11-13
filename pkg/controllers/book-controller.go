package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KarkiAnmol/BookStore-API-Golang/pkg/models"
	"github.com/KarkiAnmol/BookStore-API-Golang/pkg/utils"
	"github.com/gorilla/mux"
)

// GetBooks retrieves a list of books and sends it as a JSON response.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById retrieves details of a specific book by ID and sends it as a JSON response.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["Id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook creates a new book based on the request body and sends the created book as a JSON response.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook removes a book based on its ID and sends the deletion status as a JSON response.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deleteBook := models.DeleteBook(ID)
	res, _ := json.Marshal(deleteBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook modifies details of an existing book based on the request body and sends the updated book as a JSON response.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	// Retrieve the book details after deletion
	bookDetails, db := models.GetBookById(ID)

	// Update book details based on the request body
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save the updated book details to the database
	db.Save(&bookDetails)

	// Send the updated book details as a JSON response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
