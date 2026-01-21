package dataaccess

type UpdateTopic struct {
	ID      uint   `json:"topicID"`
	NewName string `json:"updatedName"`
}

type CreateTopic struct {
	Name string `json:"name"`
}

type DeleteTopic struct {
	ID uint `json:"topicID"`
}
