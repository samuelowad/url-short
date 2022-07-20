package models

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Url       string    `json:"url" gorm:"not null"`
	Short     string    `json:"short-link" gorm:"unique;not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
