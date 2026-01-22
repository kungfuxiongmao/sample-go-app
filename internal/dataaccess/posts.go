package dataaccess

type CreatePost struct {
	Name    string `json:"name"`
	TopicID uint   `json:"topicID"`
}

type UpdatePost struct {
	Name string `json:"updatedName"`
	ID   uint   `json:"postID"`
}

type DeletePost struct {
	ID uint `json:"postID"`
}

type FindPost struct {
	TopicID uint `json:"topicID"`
}

// Create Return Structure
type ReturnPost struct {
}
