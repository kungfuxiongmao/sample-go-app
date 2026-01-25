package dataaccess

type UpdateTopic struct {
	ID      uint   `json:"topicId"`
	NewName string `json:"updatedName"`
}

type CreateTopic struct {
	Name string `json:"name"`
}

type DeleteTopic struct {
	ID uint `json:"topicId"`
}

type GetTopic struct {
	ID uint `uri:"topicid" binding:"required"`
}
