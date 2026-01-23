package dataaccess

type CreateComment struct {
	Description string `json:"description"`
	PostID      uint   `json:"postId"`
}

type UpdateComment struct {
	Description string `json:"description"`
	ID          uint   `json:"commentId"`
}

type DeleteComment struct {
	ID uint `json:"commentId"`
}

type FindComment struct {
	PostID uint `json:"postId"`
}
