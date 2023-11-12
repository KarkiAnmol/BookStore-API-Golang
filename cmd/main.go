package main

import (
	"log"
	"net/http"

	"github.com/KarkiAnmol/BookStore-API-Golang/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8083", r))
}
