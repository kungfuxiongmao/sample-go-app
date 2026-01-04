package models

import (
	"time"
)

type Topic struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	CreatedBy uint `gorm:"not null"` //ref  users
	CreatedOn time.Time `gorm:"autoCreateTime"`
}

