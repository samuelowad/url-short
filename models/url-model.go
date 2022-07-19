package models

import "time"

type Model struct {
	ID        int       `json:"id"`
	Url       string    `json:"url"`
	Short     string    `json:"short-link"`
	createdAt time.Time `json:"-"`
	updatedAt time.Time `json:"-"`
}
