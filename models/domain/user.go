package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
