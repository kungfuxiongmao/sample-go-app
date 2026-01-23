package models

import "time"

type Comment struct {
	ID          uint      `gorm:"primaryKey" json:"commentId"`
	Description string    `gorm:"not null" json:"description"`
	User        User      `gorm:"foreignKey:CreatedBy" json:"user"`
	CreatedBy   uint      `json:"createdBy"`
	Post        Post      `gorm:"foreignKey:PostID" json:"topic"`
	PostID      uint      `json:"postId"`
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdOn"`
}
