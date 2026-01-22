package dataaccess

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
