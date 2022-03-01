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
		setRole(client, ctx, roleName)
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

func setRole(client *ent.Client, ctx context.Context, roleName string) {

	r, err := client.Role.Query().Where(role.Name(roleName)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			log.Fatalf("failed querying in database: %v", err)
		}
		builder := client.Role.Create()
		builder.SetName(roleName)
		setPermissions(builder.Mutation(), roleName)
		if _, err = builder.Save(ctx); err != nil {
			log.Fatalf("failed to create role: %v", err)
		}
	} else {
		builder := r.Update()
		resetPermissions(builder.Mutation())
		setPermissions(builder.Mutation(), roleName)
		if _, err = builder.Save(ctx); err != nil {
			log.Fatalf("failed to update permissions: %v", err)
		}
	}
}

func setPermissions(m *ent.RoleMutation, roleName string) {

	m.SetEvent("true")
	m.SetUser("true")
	switch roleName {
	case "tutor":
		m.SetUserSubscription("true")
	case "calendar":
		m.SetEventWrite("true")
	case "admin":
		m.SetUserSubscription("true")
		m.SetEventWrite("true")
		m.SetUserWrite("true")
	}
}

func resetPermissions(m *ent.RoleMutation) {
	m.SetEvent("false")
	m.SetEventWrite("false")
	m.SetUser("false")
	m.SetUserSubscription("false")
	m.SetUserWrite("false")
}
