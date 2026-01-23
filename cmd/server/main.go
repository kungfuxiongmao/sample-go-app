package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kungfuxiongmao/sample-go-app/internal/database"
	_ "github.com/kungfuxiongmao/sample-go-app/internal/init"
	"github.com/kungfuxiongmao/sample-go-app/internal/router"
)

func main() {
	//Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("JWT_KEY") == "" {
		log.Fatal("'JWT_KEY' environment variable not set!")
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
