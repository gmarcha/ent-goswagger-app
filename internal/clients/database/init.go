package database

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gamarcha/ent-goswagger-app/internal/ent"
	_ "github.com/lib/pq"
)

func Init() (client *ent.Client) {
	infos := []string{
		"host=" + os.Getenv("POSTGRES_HOST"),
		"port=" + os.Getenv("POSTGRES_PORT"),
		"user=" + os.Getenv("POSTGRES_USER"),
		"password=" + os.Getenv("POSTGRES_PASSWORD"),
		"dbname=" + os.Getenv("POSTGRES_DB"),
		"sslmode=" + os.Getenv("POSTGRES_SSLMODE"),
	}
	info := strings.Join(infos, " ")
	client, err := ent.Open("postgres", info)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	log.Println("Database connected")
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Schemas created")
	return
}