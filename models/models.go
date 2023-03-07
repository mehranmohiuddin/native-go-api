package models

type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}
