package dataaccess

import (
	"github.com/kungfuxiongmao/sample-go-app/internal/database"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAcc struct {
	Username string `json:"username"`
	Password string `json:"password"`
}