package users

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func CreateUser(c *gin.Context) {
	var u dataaccess.CreateAcc
	var user models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		api.FailMsg(c, http.StatusBadRequest, 1, err.Error())
		return
	}
	if strings.TrimSpace(u.Username) == "" || strings.TrimSpace(u.Password) == "" {
		api.FailMsg(c, http.StatusBadRequest, 3, "invalid input")
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, 2, err.Error())
		return
	}
	user.Username = u.Username
	user.Password = hashed
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, 1, "database not available")
		return
	}
	result := db.WithContext(c.Request.Context()).Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			api.FailMsg(c, http.StatusUnprocessableEntity, 3, "username already exists")
			return
		}
		api.FailMsg(c, http.StatusInternalServerError, 4, result.Error.Error())
		return
	}
	api.SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username}, "user created successfully")
	//password is not returned for obvious reasons
}

func CheckUser(c *gin.Context) {
	var u dataaccess.LoginReq
	if err := c.ShouldBindJSON(&u); err != nil {
		api.FailMsg(c, http.StatusBadRequest, 1, err.Error())
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, 1, "database not available")
		return
	}
	var user models.User
	err = db.WithContext(c.Request.Context()).Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.FailMsg(c, http.StatusUnauthorized, 1, "invalid username or password")
			return
		}
		api.FailMsg(c, http.StatusInternalServerError, 2, "database error")
		return
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(u.Password)); err != nil {
		api.FailMsg(c, http.StatusUnauthorized, 3, "invalid username or password")
		return
	}
	api.SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username}, "user authorised")
}
