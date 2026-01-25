package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

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
