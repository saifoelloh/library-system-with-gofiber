package author

import (
	"github.com/saifoelloh/minimalism_api/api/book"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name        string `gorm:"size:15" json:"name"`
	Birth       string `gorm:"type:date" json:"birth"`
	Description string `json:"description"`
	Books       []book.Book
}
