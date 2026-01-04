package main

import (
	"log"
	"github.com/kungfuxiongmao/sample-go-app/internal/router"
	"github.com/kungfuxiongmao/sample-go-app/internal/database"
)

func main() {
	db, err:=database.GetDB()
	if err != nil {
		log.Fatalf("Failed to get db: %v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	defer sqlDB.Close()
	r:=router.Setup(db)
	log.Println("Server running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
