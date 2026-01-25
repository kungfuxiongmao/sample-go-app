package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"gorm.io/gorm"
)

func DeletePost(c *gin.Context) {
	var p dataaccess.DeletePost
	var post models.Post
	//Bind input
	if err := c.ShouldBindJSON(&p); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	//Get user
	userid, exists := c.Get("userID")
	if !exists {
		api.FailMsg(c, http.StatusInternalServerError, CodeGetUserFail, "failed to retreive user ID")
		return
	}
	result := db.WithContext(c.Request.Context()).Where("id = ? AND created_by = ?", p.ID, userid).First(&post)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "failed to match record or you do not have permission")
		return
	}
	if err := deleteCommentbyPost(db.WithContext(c.Request.Context()), post.ID); err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete comments")
		return
	}
	if err := db.WithContext(c.Request.Context()).Delete(&post).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete post")
		return
	}
	api.SuccessMsg(c, nil, "successfully deleted post")
}

func deleteCommentbyPost(db *gorm.DB, postID uint) error {
	if err := db.Where("post_id = ?", postID).Delete(&models.Comment{}).Error; err != nil {
		return err
	}
	return nil
}
