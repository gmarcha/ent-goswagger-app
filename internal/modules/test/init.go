package test

import (
	"context"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
)

// Init sets database entries for test purpose.
// It is deprecated and it should be use only to generate SQL requests.
func Init(db *ent.Client) {

	createTestUsers(db)
}

func createTestUsers(db *ent.Client) {

	ctx := context.Background()

	now := time.Now()
	startAt := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())
	endAt := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+2, 0, 0, 0, now.Location())

	tutor, _ := db.Role.Query().Where(role.Name("tutor")).Only(ctx)
	calendar, _ := db.Role.Query().Where(role.Name("calendar")).Only(ctx)
	admin, _ := db.Role.Query().Where(role.Name("admin")).Only(ctx)

	_, _ = db.User.Create().SetLogin("student").Save(ctx)
	_, _ = db.User.Create().SetLogin("tutor").AddRoles(tutor).Save(ctx)
	_, _ = db.User.Create().SetLogin("calendar").AddRoles(calendar).Save(ctx)
	_, _ = db.User.Create().SetLogin("admin").AddRoles(admin).Save(ctx)

	test, _ := db.EventType.Create().SetName("test").SetColor("test").Save(ctx)
	if test != nil {
		_, _ = db.Event.Create().SetName("test").SetStartAt(startAt).SetEndAt(endAt).SetCategoryID(test.ID).Save(ctx)
	}
}
