package models

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	Text        string `json:"text"`
	Translation string `json:"translation"`
	TopicID     uint   `json:"topicId"`
	UserID      uint   `json:"userId"`
	User        User   `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
