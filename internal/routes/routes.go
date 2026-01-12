package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/users"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/users/login", users.CheckUser)
	r.POST("/users/create", users.CreateUser)
}
