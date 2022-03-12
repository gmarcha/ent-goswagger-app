package role

import (
	"context"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
	"github.com/google/uuid"
)

// Service holds an ent role client.
type Service struct {
	Role *ent.RoleClient
}

// ReadRoleByName returns a role or an error.
func (r *Service) ReadRoleByName(ctx context.Context, name string) (*ent.Role, error) {

	res, err := r.Role.Query().Where(role.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddRole returns user roles or an error.
func (r *Service) AddRole(ctx context.Context, user *user.Service, id uuid.UUID, name string) ([]*ent.Role, error) {

	userRes, roleRes, err := r.getUserAndRole(ctx, user, id, name)
	if err != nil {
		return nil, err
	}
	res, err := userRes.Update().AddRoleIDs(roleRes.ID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res.Edges.Roles, nil
}

// RemoveRole returns an error.
func (r *Service) RemoveRole(ctx context.Context, user *user.Service, id uuid.UUID, name string) error {

	userRes, roleRes, err := r.getUserAndRole(ctx, user, id, name)
	if err != nil {
		return err
	}
	_, err = userRes.Update().RemoveRoleIDs(roleRes.ID).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *Service) getUserAndRole(ctx context.Context, user *user.Service, id uuid.UUID, name string) (*ent.User, *ent.Role, error) {

	userRes, err := user.ReadUserByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	roleRes, err := r.ReadRoleByName(ctx, name)
	if err != nil {
		return nil, nil, err
	}
	return userRes, roleRes, nil
}
