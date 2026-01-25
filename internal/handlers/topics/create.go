package topics

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
