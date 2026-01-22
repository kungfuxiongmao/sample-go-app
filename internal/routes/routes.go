package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/topics"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/users"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/users/login", users.CheckUser)
	r.POST("/users/create", users.CreateUser)
	protected := r.Group("/api")
	protected.Use(middleware.RequireAuth())
	{
		protected.GET("/topics", topics.GetTopics)
		protected.POST("/topics", topics.CreateTopic)
		protected.PUT("/topics", topics.UpdateTopic)
		protected.DELETE("/topics", topics.DeleteTopic)
	}
}
