package env

import (
	"log"

	"github.com/joho/godotenv"
)

func Init(filename string) {

	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
