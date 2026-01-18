package users

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kungfuxiongmao/sample-go-app/internal/api"
	"github.com/kungfuxiongmao/sample-go-app/internal/dataaccess"
	"github.com/kungfuxiongmao/sample-go-app/internal/middleware"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	// --- 1000 SERIES: INPUT ERRORS (Client Fault) ---
	CodeBindFailed   = 1001 // JSON formatting wrong
	CodeInvalidInput = 1002 // Empty strings, validation failed
	CodeParamMissing = 1003 // URL parameters missing

	// --- 2000 SERIES: BUSINESS LOGIC ERRORS (Client Fault) ---
	CodeUserExists = 2001 // Duplicate username/email
	CodeAuthFailed = 2002 // Bad password

	// --- 5000 SERIES: SYSTEM ERRORS (Server Fault) ---
	CodeDatabaseFail = 5001 // DB connection or query failed
	CodeCryptoFail   = 5002 // Bcrypt failed
	CodeTokenGenFail = 5003 // JWT generation failed
)

func CreateUser(c *gin.Context) {
	//Initialise var
	var u dataaccess.CreateAcc
	var user models.User
	//Bind input into var
	if err := c.ShouldBindJSON(&u); err != nil {
		api.FailMsg(c, http.StatusBadRequest, CodeBindFailed, err.Error())
		return
	}
	//Validate input
	if strings.TrimSpace(u.Username) == "" || strings.TrimSpace(u.Password) == "" {
		api.FailMsg(c, http.StatusBadRequest, CodeInvalidInput, "invalid input")
		return
	}
	//Encrypt pw
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeCryptoFail, err.Error())
		return
	}
	//Initialise user model
	user.Username = u.Username
	user.Password = hashed
	db, err := middleware.GetDB(c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, "database not available")
		return
	}
	//Write into DB
	result := db.WithContext(c.Request.Context()).Create(&user)
	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				api.FailMsg(c, http.StatusUnprocessableEntity, CodeUserExists, "username already exists")
				return
			}
		}
		api.FailMsg(c, http.StatusInternalServerError, CodeDatabaseFail, result.Error.Error())
		return
	}
	//Create cookie to maintain login
	tokenString, err := middleware.CreateToken(user.ID, c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeTokenGenFail, "Session creation failed")
		return
	}
	api.SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username, "token": tokenString}, "user created successfully")
}

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

	tokenString, err := middleware.CreateToken(user.ID, c)
	if err != nil {
		api.FailMsg(c, http.StatusInternalServerError, CodeTokenGenFail, "Session creation failed")
		return
	}
	api.SuccessMsg(c, gin.H{"id": user.ID, "username": user.Username, "token": tokenString}, "user authorised")
}
