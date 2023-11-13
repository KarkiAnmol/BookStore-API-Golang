// Package config provides functionality for configuring and connecting to the database.
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect initializes the database connection using the gorm library.
// It connects to a MySQL database with the specified connection string.
// If an error occurs during the connection, it panics.
func Connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the instance of the connected database.
func GetDB() *gorm.DB {
	return db
}
