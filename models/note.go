package models

import "gorm.io/gorm"

// Book model
type Note struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateNote struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetNotes []struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
