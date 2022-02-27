package database

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/migrate"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
	_ "github.com/lib/pq"
)

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

	for _, roleName := range []string{"tutor", "calendar", "admin"} {
		createRole(client, ctx, roleName)
	}
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

func createRole(client *ent.Client, ctx context.Context, roleName string) {

	_, err := client.Role.Query().Where(role.Name(roleName)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			log.Fatalf("failed querying in database: %v", err)
		}
		builder := client.Role.Create()
		builder.SetName(roleName)
		createPermissions(builder.Mutation(), roleName)
		if _, err = builder.Save(ctx); err != nil {
			log.Fatalf("failed to create role: %v", roleName)
		}
	}
}

func createPermissions(m *ent.RoleMutation, roleName string) {

	switch roleName {
	case "tutor":
		m.SetEvent(true)
		m.SetUser(true)
		m.SetUserSubscription(true)
	case "calendar":
		m.SetEventWrite(true)
	case "admin":
		m.SetUserWrite(true)
	}
}
