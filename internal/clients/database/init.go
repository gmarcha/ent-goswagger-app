package database

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	_ "github.com/lib/pq"
)

func Init() (client *ent.Client) {

	database := "postgres"
	infos := []string{
		"host=" + os.Getenv("POSTGRES_HOST"),
		"port=" + os.Getenv("POSTGRES_PORT"),
		"user=" + os.Getenv("POSTGRES_USER"),
		"password=" + os.Getenv("POSTGRES_PASSWORD"),
		"dbname=" + os.Getenv("POSTGRES_DB"),
		"sslmode=" + os.Getenv("POSTGRES_SSLMODE"),
	}
	info := strings.Join(infos, " ")

	client, err := ent.Open(database, info)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	log.Println("Database connected")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Schemas created")
	return
}
