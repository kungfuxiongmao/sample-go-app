package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

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
