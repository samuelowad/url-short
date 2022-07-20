package database

import (
	models "github.com/samuelowad/url-shorterner/models"
)

func FindUrl(url string) (models.Model, error) {
	var data models.Model

	err := database.Where("short=?", url).Find(&data)
	if err != nil {
		return data, err.Error
	}

	return data, nil

}

func CreateUrl(data *models.Model) error {

	err := database.Create(&data).Error

	return err

}
