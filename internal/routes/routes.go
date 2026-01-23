package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/comments"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/posts"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/topics"
	"github.com/kungfuxiongmao/sample-go-app/internal/handlers/users"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/users/login", users.CheckUser)
	r.POST("/users/create", users.CreateUser)
	r.POST("/logout", users.Logout)
	protected := r.Group("/api")
	protected.Use(middleware.RequireAuth())
	{
		protected.GET("/me", users.GetProfile)
		protected.GET("/topics", topics.GetTopics)
		protected.POST("/topics", topics.CreateTopic)
		protected.PUT("/topics", topics.UpdateTopic)
		protected.DELETE("/topics", topics.DeleteTopic)
		protected.GET("/topics/:topicid", posts.GetPost)

		protected.POST("/posts", posts.CreatePost)
		protected.PUT("/posts", posts.UpdatePost)
		protected.DELETE("/posts", posts.DeletePost)
		protected.GET("posts/:postid", comments.GetComments)

		protected.POST("/comments", comments.CreateComment)
		protected.PUT("/comments", comments.UpdateComment)
		protected.DELETE("/comments", comments.DeleteComment)
	}
}
