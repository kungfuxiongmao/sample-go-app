package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBToContext(db *gorm.DB) gin.HandlerFunc {
	if db == nil {
		return nil
	}
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func GetDB(c *gin.Context) (*gorm.DB, error) {
	extr, exists := c.Get("db")
	if !exists {
		return nil, fmt.Errorf("Database does not exist")
	}
	db, ok := extr.(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("Error in getting database")
	}
	return db, nil
}
