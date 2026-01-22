package dataaccess

type CreateComment struct {
	Description string `json:"description"`
	PostID      uint   `json:"postID"`
}

type UpdateComment struct {
	Description string `json:"description"`
	ID          uint   `json:"commentID"`
}

type DeleteComment struct {
	ID uint `json:"commentID"`
}

type FindComment struct {
	PostID uint `json:"postID"`
}
