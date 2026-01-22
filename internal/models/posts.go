package models

import "time"

type Post struct {
	ID          uint   `gorm:"primaryKey" json:"postid"`
	PostName    string `gorm:"not null" json:"postname"`
	Description string `gorm:"not null" json:"description"`
	User        User   `gorm:"foreignKey:CreatedBy"`
	CreatedBy   uint
	Topic       Topic `gorm:"foreignKey:TopicID" json:"omitempty"`
	TopicID     uint
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdon"`
}
