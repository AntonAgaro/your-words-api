package models

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Words       []Word `gorm:"foreignKey:TopicID"`
}
