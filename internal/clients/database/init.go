package database

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/migrate"
	// Import postgres dependencies.
	_ "github.com/lib/pq"
)

// Init returns an ent client.
func Init() *ent.Client {

	database := "postgres"
	info := createPostgresInfo()

	client, err := ent.Open(database, info)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	log.Println("Database connected")

	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Schemas created")

	return client
}

func createPostgresInfo() string {

	infos := []string{
		"host=" + os.Getenv("POSTGRES_HOST"),
		"port=" + os.Getenv("POSTGRES_PORT"),
		"user=" + os.Getenv("POSTGRES_USER"),
		"password=" + os.Getenv("POSTGRES_PASSWORD"),
		"dbname=" + os.Getenv("POSTGRES_DB"),
		"sslmode=" + os.Getenv("POSTGRES_SSLMODE"),
	}
	return strings.Join(infos, " ")
}
