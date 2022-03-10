package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {

	env := os.Args[1]
	filename := "./config/.env." + env
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading environment file")
	}
}
