package database

import (
	"log"

	models "github.com/samuelowad/url-shorterner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://postgres:postgres@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Model{})

	return db
}
