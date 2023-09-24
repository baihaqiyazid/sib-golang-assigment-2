package app

import (
	// "assignment-2/models"
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "gosibtwo"
	db       *gorm.DB
	err      error
	// dns      = "postgres://postgres:postgres@localhost:5432/goreactmovies?sslmode=disable"
)

func StartDB()  {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil{
		log.Fatal("error connecting ti database:", err)
	}

	db.Debug().AutoMigrate( models.Order{}, models.Item{})
}

func GetDB() *gorm.DB{
	return db
}