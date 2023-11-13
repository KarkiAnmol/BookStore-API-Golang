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

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["Id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deleteBook := models.DeleteBook(ID)
	res, _ := json.Marshal(deleteBook)
	w.Header().Set("Contet-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
