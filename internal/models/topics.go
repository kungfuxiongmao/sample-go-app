package models

import "time"

type Topics struct {
	ID          uint   `gorm:"primaryKey"`
	TopicName   string `gorm:"not null"`
	User        User   `gorm:"foreignKey:CreatedBy"`
	CreatedBy   uint
	DateCreated time.Time `gorm:"autoCreateTime"`
}
