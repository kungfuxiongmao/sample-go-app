package models

import "time"

type Topic struct {
	ID          uint      `gorm:"primaryKey" json:"topicId"`
	TopicName   string    `gorm:"not null" json:"topicName"`
	User        User      `gorm:"foreignKey:CreatedBy"`
	CreatedBy   uint      `json:"createdBy"`
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdOn"`
}
