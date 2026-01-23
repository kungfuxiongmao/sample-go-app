package models

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"postId"`
	PostName    string    `gorm:"not null" json:"postName"`
	Description string    `gorm:"not null" json:"description"`
	User        User      `gorm:"foreignKey:CreatedBy" json:"user"`
	CreatedBy   uint      `json:"createdBy"`
	Topic       Topic     `gorm:"foreignKey:TopicID" json:"topic"`
	TopicID     uint      `json:"topicId"`
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdOn"`
}
