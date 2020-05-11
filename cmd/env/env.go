package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Get - get env value
func Get(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	return os.Getenv(key)
}
