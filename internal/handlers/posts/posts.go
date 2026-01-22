package posts

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

func CreatePost(c *gin.Context) {
	var p dataaccess.CreatePost
	var post models.Post
	var topic models.Topic
	//Bind input into var
	if err := c.ShouldBindJSON(&p); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	//Validate Input
	if p.Name == "" || p.Description == "" {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "empty title or description")
		return
	}
	//Get User
	userid, exists := c.Get("userID")
	if !exists {
		api.FailMsg(c, http.StatusInternalServerError, CodeGetUserFail, "failed to retreive user ID")
		return
	}
	//Check that topicID is valid
	//Get DB
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	result := db.Select("id").First(&topic, "id = ?", p.TopicID)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "topic not found")
		return
	}
	//Update post
	post.CreatedBy = userid.(uint)
	post.PostName = p.Name
	post.Description = p.Description
	post.TopicID = topic.ID
	//Update DB
	result = db.WithContext(c.Request.Context()).Create(&post)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, post, "post created")
}

func GetPost(c *gin.Context) {
	var p dataaccess.FindPost
	var post []models.Post
	//Bind input into var
	if err := c.ShouldBindJSON(&p); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	result := db.WithContext(c.Request.Context()).Preload("User").Where("topic_id = ?", p.TopicID).Find(&post)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, post, "successfully received posts")
}

func UpdatePost(c *gin.Context) {
	var p dataaccess.UpdatePost
	var post models.Post
	//Bind input
	if err := c.ShouldBindJSON(&p); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	if p.Name == "" || p.Description == "" {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "empty title or description")
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
	// Only authorised posts will be taken
	result := db.WithContext(c.Request.Context()).Where("id = ? AND created_by = ?", p.ID, userid).First(&post)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "topic not found or you do not have permission")
		return
	}
	post.PostName = p.Name
	post.Description = p.Description
	if err := db.WithContext(c.Request.Context()).Save(&post).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, err.Error())
		return
	}
	api.SuccessMsg(c, post, "updated topic name")
}

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
	if err := db.WithContext(c.Request.Context()).Delete(&post).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete")
		return
	}
	api.SuccessMsg(c, nil, "successfully deleted post")
}
