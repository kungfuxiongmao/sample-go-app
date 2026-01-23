package database

import (
	"fmt"
	"os"

	"github.com/kungfuxiongmao/sample-go-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnector() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")
	return fmt.Sprintf("host=%s port=%v sslmode=%s user=%s password=%s dbname=%s",
		host, port, ssl, user, password, dbname)
}

func GetDB() (*gorm.DB, error) {
	dsn := GetConnector()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open GORM DB: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from GORM DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Topic{}, &models.Post{}, &models.Comment{}) //include automigrate
	if err != nil {
		return nil, err
	}
	return db, nil
}
