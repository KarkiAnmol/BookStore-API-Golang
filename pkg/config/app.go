// Package config provides functionality for configuring and connecting to the database.
package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

// Connect initializes the database connection using the gorm library.
// It connects to a MySQL database with the specified connection string.
// If an error occurs during the connection, it panics.
func Connect() {
	d, err := gorm.Open("mysql", "anmol-Karki@123/simpleREST?charset=UTF8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the instance of the connected database.
func GetDB() *gorm.DB {
	return db
}
