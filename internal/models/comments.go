package models

import {
	"gorm.io/gorm"
}

type Comments struct {
	ID int `gorm: "PrimaryKey"`
	PostID int `gorm: `
	ParentID int `gorm: `
	Name string `gorm: not null`
	CreatedBy int `gorm` //ref  users
	CreatedOn datetime
}

type CommentsLikes struct {
	CommentID int 
	UserID int
	Key (CommentID, UserID) `gorm: PrimaryKey`
}