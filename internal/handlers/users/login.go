package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	CodeBindFailed   = 1001 // JSON formatting wrong
	CodeInvalidInput = 1002
	CodeParamMissing = 1003 // URL parameters missing

	CodeUserExists = 2001 // Duplicate username/email
	CodeAuthFailed = 2002

	CodeDatabaseFail = 5001 // DB connection or query failed
	CodeCryptoFail   = 5002
	CodeTokenGenFail = 5003
	CodeGetUserFail  = 5004
)

func CheckUser(c *gin.Context) {
	var u dataaccess.LoginReq
	if err := c.ShouldBindJSON(&u); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	var user models.User
	err = db.WithContext(c.Request.Context()).Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.FailMsg(c, http.StatusUnauthorized, CodeAuthFailed, "invalid username or password")
			return
		}
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database error")
		return
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(u.Password)); err != nil {
		api.FailMsg(c, http.StatusUnauthorized, CodeAuthFailed, "invalid username or password")
		return
	}

	tokenString, err := middleware.CreateToken(user.ID, user.Username, c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeTokenGenFail, "Session creation failed")
		return
	}
	api.SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username, "token": tokenString}, "user authorised")
}
