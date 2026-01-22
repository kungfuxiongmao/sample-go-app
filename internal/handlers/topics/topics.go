package topics

//Add Context for db queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

const (
	CodeBindFailed   = 1001
	CodeInvalidInput = 1002

	CodeDatabaseFail = 5001
	CodeGetUserFail  = 5004
)

func CreateTopic(c *gin.Context) {
	//Initialise variables
	var t dataaccess.CreateTopic
	var topic models.Topic
	//Bind input
	if err := c.ShouldBindJSON(&t); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	if t.Name == "" {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "empty title")
		return
	}
	//Get user
	userid, exists := c.Get("userID")
	if !exists {
		api.FailMsg(c, http.StatusInternalServerError, CodeGetUserFail, "failed to retreive user ID")
		return
	}
	topic.CreatedBy = userid.(uint)
	topic.TopicName = t.Name
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	//Write into DB
	tx := db.Begin()
	if err := tx.Create(&topic).Error; err != nil {
		tx.Rollback()
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, err.Error())
		return
	}
	tx.Commit()
	api.SuccessMsg(c, gin.H{"id": topic.ID, "name": topic.TopicName}, "created topic successfully")
}

func GetTopics(c *gin.Context) {
	var topics []models.Topic
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	result := db.Preload("User").Find(&topics)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, topics)
}

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
	if err := db.WithContext(c.Request.Context()).Delete(&topic).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete")
		return
	}
	api.SuccessMsg(c, nil, "successfully deleted topic")
}
