// Package routes defines the routes for the BookStore API using the gorilla/mux router.
package routes

import (
	"github.com/KarkiAnmol/BookStore-API-Golang/pkg/controllers" // Importing the controllers package for route handling
	"github.com/gorilla/mux"                                     // Importing the gorilla/mux package for router implementation
)

// RegisterBookStoreRoutes is a function that registers routes for the BookStore API.
// It takes a router (mux.Router) as a parameter and defines various routes with their corresponding controller functions.
var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Define route for creating a new book (HTTP POST)
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	// Define route for retrieving all books (HTTP GET)
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")

	// Define route for retrieving a specific book by ID (HTTP GET)
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")

	// Define route for updating a book by ID (HTTP PUT)
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")

	// Define route for deleting a book by ID (HTTP DELETE)
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
