package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
)

func Logout(c *gin.Context) {
	middleware.ClearToken(c)
	api.SuccessMsg(c, nil, "successfully logged out")
}
