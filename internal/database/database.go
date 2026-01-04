package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)


func GetConnector() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")
	return fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s", 
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
	db.AutoMigrate(&User{}) //include automigrate
	return db, nil
}

