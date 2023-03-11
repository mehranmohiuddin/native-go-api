package models

type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Year     int    `json:"year" validate:"required"`
	Director string `json:"director" validate:"required"`
}
