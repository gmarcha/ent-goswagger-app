package role

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
)

func Init(api *operations.TutorAPI, db *ent.Client) {

	userService := &user.Service{User: db.User}
	roleService := &Service{Role: db.Role}

	api.RoleAddTutorHandler = &addTutor{user: userService, role: roleService}
	api.RoleRemoveTutorHandler = &removeTutor{user: userService, role: roleService}

	api.RoleAddCalendarHandler = &addCalendar{user: userService, role: roleService}
	api.RoleRemoveCalendarHandler = &removeCalendar{user: userService, role: roleService}

	api.RoleAddAdminHandler = &addAdmin{user: userService, role: roleService}
	api.RoleRemoveAdminHandler = &removeAdmin{user: userService, role: roleService}
}
