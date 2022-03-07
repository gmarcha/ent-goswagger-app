package role

import (
	"context"
	"log"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
)

func Init(api *operations.TutorAPI, db *ent.Client) {

	userService := &user.Service{User: db.User}
	roleService := &Service{Role: db.Role}

	createRoles(db)

	api.RoleAddTutorHandler = &addTutor{user: userService, role: roleService}
	api.RoleRemoveTutorHandler = &removeTutor{user: userService, role: roleService}

	api.RoleAddCalendarHandler = &addCalendar{user: userService, role: roleService}
	api.RoleRemoveCalendarHandler = &removeCalendar{user: userService, role: roleService}

	api.RoleAddAdminHandler = &addAdmin{user: userService, role: roleService}
	api.RoleRemoveAdminHandler = &removeAdmin{user: userService, role: roleService}
}

func createRoles(db *ent.Client) {

	ctx := context.Background()
	for _, roleName := range []string{"tutor", "calendar", "admin"} {
		setRole(db, ctx, roleName)
	}
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
