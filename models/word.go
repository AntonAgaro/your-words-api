package models

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	Text        string `json:"text"`
	Translation string `json:"translation"`
	TopicID     uint   `json:"topicId"`
}
