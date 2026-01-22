package models

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"postid"`
	PostName    string    `gorm:"not null" json:"postname"`
	Description string    `gorm:"not null" json:"description"`
	User        User      `gorm:"foreignKey:CreatedBy" json:"use"`
	CreatedBy   uint      `json:"createdby"`
	Topic       Topic     `gorm:"foreignKey:TopicID" json:"topic"`
	TopicID     uint      `json:"topicid"`
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdon"`
}
