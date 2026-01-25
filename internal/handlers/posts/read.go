package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

func GetPost(c *gin.Context) {
	var p dataaccess.GetPost
	var post []models.Post
	if err := c.ShouldBindUri(&p); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, err.Error())
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	result := db.Debug().WithContext(c.Request.Context()).Where("topic_id = ?", p.TopicID).Preload("Topic").Preload("User").Find(&post)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, post, "successfully received posts")
}

func FindPost(c *gin.Context) {
	var p dataaccess.FindPost
	var post models.Post
	if err := c.ShouldBindUri(&p); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "invalid post ID")
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	result := db.WithContext(c.Request.Context()).Where("id = ?", p.PostID).Preload("User").First(&post)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, post, "successfully received post")
}
