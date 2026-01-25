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
	r.GET("/topic/:topicid", topics.FindTopic)
	r.GET("/topics", topics.GetTopics)
	r.GET("/topics/:topicid", posts.GetPost)
	r.GET("/post/:postid", posts.FindPost)
	r.GET("/posts/:postid", comments.GetComments)
	r.GET("/comment/:commentid", comments.FindComment)

	protected := r.Group("/api")
	protected.Use(middleware.RequireAuth())
	{
		protected.GET("/me", users.GetProfile)
		protected.POST("/topics", topics.CreateTopic)
		protected.PUT("/topics", topics.UpdateTopic)
		protected.DELETE("/topics", topics.DeleteTopic)

		protected.POST("/posts", posts.CreatePost)
		protected.PUT("/posts", posts.UpdatePost)
		protected.DELETE("/posts", posts.DeletePost)

		protected.POST("/comments", comments.CreateComment)
		protected.PUT("/comments", comments.UpdateComment)
		protected.DELETE("/comments", comments.DeleteComment)
	}
}
