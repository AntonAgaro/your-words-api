package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string
	Email    string `json:"email"`
	Words    []Word
}
