package init

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("JWT_KEY") == "" {
		log.Fatal("'JWT_KEY' environment variable not set!")
	}
}
