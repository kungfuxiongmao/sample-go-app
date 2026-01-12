package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kungfuxiongmao/sample-go-app/internal/database"
	"github.com/kungfuxiongmao/sample-go-app/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("Failed to get db: %v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	defer sqlDB.Close()
	r := router.Setup(db)
	log.Println("Server running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
