package users

import (
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(c *gin.Context) {
	var u dataaccess.CreateAcc
	var user models.User
	if err := c.ShouldBindJSON(&u); err != nil {
    	FailMsg(c, http.StatusBadRequest, 1, err)
		return
    }
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost,)
	if err != nil {
		FailMsg(c, http.StatusInternalServerError, 2, err)
		return
	}
	user.Username = u.Username
	user.Password = hashed
	db := middleware.GetDB(c)
	ctx := c.Request.Context()
	err = gorm.G[models.User](db).Create(ctx, &user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			FailMsg(c, http.StatusConflict, 3, "username already exists")
			return
		}
		FailMsg(c, http.StatusInternalServerError, 4, "database error")
		return
	}
	SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username,}, "user created successfully")
	//password is not returned for obvious reasons
	return 
}

func CheckUser(c *gin.Context) {
	var u dataaccess.LoginReq
	if err := c.ShouldBindJSON(&u); err != nil {
    	FailMsg(c, http.StatusBadRequest, 1, err)
		return
    }
	db := middleware.GetDB(c)
	user, err := gorm.G[models.User](db).Where("username = ?", u.Username).First(c.Request.Context())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			FailMsg(c, http.StatusUnauthorized, 1, "invalid username or password")
			return
		}
		FailMsg(c, http.StatusInternalServerError, 2, "database error")
		return
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(u.Password)); err != nil {
		FailMsg(c, http.StatusUnauthorized, 3, "invalid username or password")
		return
	}
	SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username,}, "user authorised")
	return 
}