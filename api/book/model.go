package book

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Release     string `gorm:"type:date" json:"release"`
	AuthorID    int    `json:"author_id"`
}
