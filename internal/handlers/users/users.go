package users

import (
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"net/http"

	
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/pkg/errors"
)

const (
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func CreateUser(c *gin.Context) (int, error) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
    	FailMsg(c, http.StatusInternalServerError, 1, err)
		return
    }
	hashed, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost,)
	if err != nil {
		FailMsg(c, http.StatusInternalServerError, 2, err)
		return
	}
	db := middleware.GetDB(c)
	//get id and write into db, return id
}

func CheckUser()