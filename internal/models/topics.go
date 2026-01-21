package models

import "time"

type Topic struct {
	ID          uint   `gorm:"primaryKey" json:"topicid"`
	TopicName   string `gorm:"not null" json:"topicname"`
	User        User   `gorm:"foreignKey:CreatedBy"`
	CreatedBy   uint
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdon"`
}
