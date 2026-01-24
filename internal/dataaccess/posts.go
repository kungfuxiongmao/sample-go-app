package dataaccess

type CreatePost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TopicID     uint   `json:"topicId"`
}

type UpdatePost struct {
	Name        string `json:"updatedName"`
	Description string `json:"description"`
	ID          uint   `json:"postId"`
}

type DeletePost struct {
	ID uint `json:"postId"`
}

type FindPost struct {
	TopicID uint `uri:"topicid" binding:"required"`
}
