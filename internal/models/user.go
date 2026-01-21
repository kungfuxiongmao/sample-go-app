package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"userid"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password []byte `gorm:"not null" json:"-"` //hashed password
}
