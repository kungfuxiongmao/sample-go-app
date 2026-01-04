package models

type User struct {
	ID   uint 	`gorm:"primaryKey"`
	Username string 	`gorm:"uniqueIndex;not null"`
	Password []byte 	`gorm:"not null"` //hashed password
}
