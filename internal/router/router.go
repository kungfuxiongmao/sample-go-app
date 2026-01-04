package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/routes"
	"gorm.io/gorm"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
)

func Setup(db *gorm.DB) *gin.Engine {
	r:= gin.Default()
	r.Use(middleware.DBToContext(db))
	routes.GetRoutes(r)
	return r
}