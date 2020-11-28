package models

import (
	"time"

	"gorm.io/gorm"
)

// Account model
type Account struct {
	gorm.Model
	Username  string
	FullName  string
	Password  string
	Email     string
	DateBirth time.Time `json:"date_birth"`
	AdminRole bool      `gorm:"default:0"`
}
