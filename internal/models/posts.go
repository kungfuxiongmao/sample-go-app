package models

import {
	"gorm.io/gorm"
}

type Posts struct {
	ID int `gorm: "PrimaryKey"`
	TopicID int `gorm:`
	Name string `gorm: not null`
	Description string `gorm: not null`
	CreatedBy int `gorm` //ref  users
	CreatedOn datetime
}

