package dataaccess

import (
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAcc struct {
	Username string `json:"username"`
	Password string `json:"password"`
}