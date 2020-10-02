package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name, Description string
	Email             *string
	Birth             *time.Time
}
