package bookstorage

import "time"

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	YearPublished int       `json:"year_published"`
	AddedAt       time.Time `json:"added_at"`
}
