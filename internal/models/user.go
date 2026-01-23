package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"userId"`
	Username string `gorm:"uniqueIndex;not null" json:"userName"`
	Password []byte `gorm:"not null" json:"-"` //hashed password
}
