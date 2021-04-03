package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get : get env value
func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Loading .env File")
	}
	return os.Getenv(key)
}
