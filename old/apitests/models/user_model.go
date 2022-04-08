package models

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Username string
	Token    string
}

// Profile ...
type Profile struct {
	gorm.Model
	User User
	Type string
}
