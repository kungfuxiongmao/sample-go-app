package comments

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

func CreateComment(c *gin.Context) {
	var cm dataaccess.CreateComment
	var comment models.Comment
	var post models.Post
	//Bind input into var
	if err := c.ShouldBindJSON(&cm); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	//Validate Input
	if cm.Description == "" {
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
	result := db.Select("id").First(&post, "id = ?", cm.PostID)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "post not found")
		return
	}
	//Update Comment
	comment.CreatedBy = userid.(uint)
	comment.Description = cm.Description
	comment.PostID = cm.PostID
	//Update DB
	result = db.WithContext(c.Request.Context()).Create(&comment)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, comment, "post created")
}

func GetComments(c *gin.Context) {
	var cm dataaccess.FindComment
	var comments []models.Comment
	//Bind input into var
	if err := c.ShouldBindJSON(&cm); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	//Get DB
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	//Search DB to return comments with Post info and User info
	result := db.WithContext(c.Request.Context()).Preload("User").Preload("Post").Where("post_id = ?", cm.PostID).Find(&comments)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, comments, "successfully received posts")
}

func UpdateComment(c *gin.Context) {
	var cm dataaccess.UpdateComment
	var comment models.Comment
	//Bind input
	if err := c.ShouldBindJSON(&cm); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	//Validate Input
	if cm.Description == "" {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "empty title or description")
		return
	}
	//Get DB
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
	// Only authorised comments will be taken
	result := db.WithContext(c.Request.Context()).Where("id = ? AND created_by = ?", cm.ID, userid).First(&comment)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "topic not found or you do not have permission")
		return
	}
	comment.Description = cm.Description
	if err := db.WithContext(c.Request.Context()).Save(&comment).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, err.Error())
		return
	}
	api.SuccessMsg(c, comment, "updated comment name")
}

func DeleteComment(c *gin.Context) {
	var cm dataaccess.DeleteComment
	var comment models.Comment
	//Bind input
	if err := c.ShouldBindJSON(&cm); err != nil {
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
	result := db.WithContext(c.Request.Context()).Where("id = ? AND created_by = ?", cm.ID, userid).First(&comment)
	if result.Error != nil {
		api.FailMsg(c, http.StatusNotFound, CodeDatabaseFail, "failed to match record or you do not have permission")
		return
	}
	if err := db.WithContext(c.Request.Context()).Delete(&comment).Error; err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "failed to delete")
		return
	}
	api.SuccessMsg(c, nil, "successfully deleted comment")
}
