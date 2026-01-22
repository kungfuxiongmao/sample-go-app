package dataaccess

import (
	"time"

	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

type CreatePost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TopicID     uint   `json:"topicID"`
}

type UpdatePost struct {
	Name        string `json:"updatedName"`
	Description string `json:"description"`
	ID          uint   `json:"postID"`
}

type DeletePost struct {
	ID uint `json:"postID"`
}

type FindPost struct {
	TopicID uint `json:"topicID"`
}

// Create Return Structure
type PostResponse struct { //for GET Methods
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ID          uint        `json:"postID"`
	User        models.User `json:"user"`
	CreatedBy   uint
	TopicID     uint
	DateCreated time.Time `gorm:"autoCreateTime" json:"createdon"`
}
}
