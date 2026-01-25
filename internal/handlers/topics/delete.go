package topics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"gorm.io/gorm"
)

func DeleteTopic(c *gin.Context) {
	var t dataaccess.DeleteTopic
	var topic models.Topic
	//Bind input
	if err := c.ShouldBindJSON(&t); err != nil {
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
	result := db.WithContext(c.Request.Context()).Where("id = ? AND created_by = ?", t.ID, userid).First(&topic)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "failed to match record or you do not have permission")
		return
	}
	if err := deletePostsByTopicID(db.WithContext(c.Request.Context()), topic.ID); err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete posts/comments")
		return
	}
	if err := db.WithContext(c.Request.Context()).Delete(&topic).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete topic")
		return
	}
	api.SuccessMsg(c, nil, "successfully deleted topic")
}

func deletePostsByTopicID(db *gorm.DB, topicID uint) error {
	var posts []models.Post
	if err := db.Where("topic_id = ?", topicID).Find(&posts).Error; err != nil {
		return err
	}
	for _, post := range posts {
		if err := db.Where("post_id = ?", post.ID).Delete(&models.Comment{}).Error; err != nil {
			return err
		}
		if err := db.Delete(&post).Error; err != nil {
			return err
		}
	}
	return nil
}
