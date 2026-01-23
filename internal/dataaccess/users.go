package dataaccess

type LoginReq struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type CreateAcc struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}
