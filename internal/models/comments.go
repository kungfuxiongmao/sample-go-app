package models

import "time"

type Comment struct {
	ID          uint      `gorm:"primaryKey" json:"commentid"`
	Description string    `gorm:"not null" json:"description"`
	User        User      `gorm:"foreignKey:CreatedBy" json:"use"`
	CreatedBy   uint      `json:"createdby"`
	Post        Post      `gorm:"foreignKey:PostID" json:"topic"`
	PostID      uint      `json:"postid"`
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdon"`
}
