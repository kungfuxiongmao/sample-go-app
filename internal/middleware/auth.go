package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	isProd := os.Getenv("GO_ENV") == "production"
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 24*10*3600, "/", "", isProd, true)

	return tokenString, nil
}

func ClearToken(c *gin.Context) {
	isProd := os.Getenv("GO_ENV") == "production"
	c.SetCookie("Auth", "", -1, "/", "", isProd, true)
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the cookie off the request
		tokenString, err := c.Cookie("Auth")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie missing"})
			return
		}

		// 2. Parse and Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
			// Security Check: Ensure the alg is HMAC (HS256)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		// 3. Check for parsing errors (Expired, fake signature, etc)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 4. Extract the Claims to get the UserID
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
			userid, err := strconv.Atoi(claims.Subject)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID"})
				return
			}
			//Stores UserID as an uint to use for DB
			c.Set("userID", uint(userid))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}
