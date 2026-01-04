package dataaccess

import (
	"github.com/kungfuxiongmao/sample-go-app/internal/database"
	"github.com/kungfuxiongmao/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}
