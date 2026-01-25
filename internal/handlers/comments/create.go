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
	result := db.Select("id").Where("id=?", cm.PostID).First(&post)
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
