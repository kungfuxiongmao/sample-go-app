package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint, c *gin.Context) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "cvwo_assign",
		Subject:   fmt.Sprint(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 10 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("jwtKey")))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 24*10*3600, "/", "", false, false)
	return tokenString, err
}

func CheckToken(token string) (jwt.RegisteredClaims, error) {
	claim := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(token, *claim)
}
