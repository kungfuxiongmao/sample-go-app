package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

func GetComments(c *gin.Context) {
	var cm dataaccess.GetComment
	var comments []models.Comment
	//Bind input into var
	if err := c.ShouldBindUri(&cm); err != nil {
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
	result := db.WithContext(c.Request.Context()).Preload("User").Preload("Post").Preload("Post.User").Where("post_id = ?", cm.PostID).Find(&comments)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, comments, "successfully received posts")
}

func FindComment(c *gin.Context) {
	var cm dataaccess.FindComment
	var comment models.Comment
	if err := c.ShouldBindUri(&cm); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "invalid comment ID")
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	result := db.WithContext(c.Request.Context()).Where("id = ?", cm.CommentID).Preload("User").First(&comment)
	if result.Error != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	api.SuccessMsg(c, comment, "successfully received post")
}
