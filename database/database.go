package database

import (
	"log"

	models "github.com/samuelowad/url-shorterner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	dbUrl := "postgres://postgres:postgres@localhost:5432/url"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	database = db
	err = db.AutoMigrate(&models.Model{})
	if err != nil {
		log.Fatal(err)
	}

}
