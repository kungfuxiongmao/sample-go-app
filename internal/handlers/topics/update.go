package topics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

func UpdateTopic(c *gin.Context) {
	var t dataaccess.UpdateTopic
	var topic models.Topic
	//Bind input
	if err := c.ShouldBindJSON(&t); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	if t.NewName == "" {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "empty title")
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
	// Only authorised topics will be taken
	result := db.WithContext(c.Request.Context()).Where("id = ? AND created_by = ?", t.ID, userid).First(&topic)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "topic not found or you do not have permission")
		return
	}
	topic.TopicName = t.NewName
	if err := db.Save(&topic).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, err.Error())
		return
	}
	api.SuccessMsg(c, gin.H{"id": topic.ID, "name": topic.TopicName}, "updated topic name")
}
