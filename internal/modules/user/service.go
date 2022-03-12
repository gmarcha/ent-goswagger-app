package user

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/user"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// Service holds an ent user client.
type Service struct {
	User *ent.UserClient
}

// ListUser returns a user list or an error.
func (s *Service) ListUser(ctx context.Context, tutor *bool) ([]*ent.User, error) {

	builder := s.User.Query()
	if tutor != nil && *tutor {
		builder.WithRoles(func(q *ent.RoleQuery) {
			q.Where(role.Name("tutor"))
		})
	}
	res, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser returns a new created user or an error.
func (s *Service) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {

	builder := s.User.Create().SetLogin(user.Login)
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ReadUserByLogin returns a user or an error.
func (s *Service) ReadUserByLogin(ctx context.Context, login string) (*ent.User, error) {

	res, err := s.User.Query().Where(user.Login(login)).WithRoles().WithEvents().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUserByLogin returns an existing updated user or an error.
func (s *Service) UpdateUserByLogin(ctx context.Context, login string, userInfo *ent.User) (*ent.User, error) {

	user, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	builder := s.User.UpdateOne(user)
	setUser(builder.Mutation(), userInfo)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteUserByLogin returns an error.
func (s *Service) DeleteUserByLogin(ctx context.Context, login string) error {

	_, err := s.User.Delete().Where(user.Login(login)).Exec(ctx)
	return err
}

// ListUserEventsByLogin returns an event list of a user or an error.
func (s *Service) ListUserEventsByLogin(ctx context.Context, login string) ([]*ent.Event, error) {

	res, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, nil
}

// ListUserRolesByLogin returns a role list of a user or an error.
func (s *Service) ListUserRolesByLogin(ctx context.Context, login string) ([]*ent.Role, error) {

	res, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return res.Edges.Roles, nil
}

// SubscribeUserByLogin returns an event list of a user or an error.
func (s *Service) SubscribeUserByLogin(ctx context.Context, login string, id uuid.UUID) ([]*ent.Event, error) {

	user, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	res, err := user.Update().AddEventIDs(id).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, nil
}

// UnsubscribeUserByLogin returns an error.
func (s *Service) UnsubscribeUserByLogin(ctx context.Context, login string, id uuid.UUID) error {

	user, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return err
	}
	_, err = user.Update().RemoveEventIDs(id).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// ReadUserByID returns a user or an error.
func (s *Service) ReadUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {

	res, err := s.User.Query().Where(user.ID(id)).WithRoles().WithEvents().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUserByID returns an existing updated user or an error.
func (s *Service) UpdateUserByID(ctx context.Context, id uuid.UUID, user *ent.User) (*ent.User, error) {

	builder := s.User.UpdateOneID(id)
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteUserByID returns an error.
func (s *Service) DeleteUserByID(ctx context.Context, id uuid.UUID) error {

	return s.User.DeleteOneID(id).Exec(ctx)
}

// ListUserEventsByID returns an event list of a user or an error.
func (s *Service) ListUserEventsByID(ctx context.Context, id uuid.UUID) ([]*ent.Event, error) {

	res, err := s.ReadUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, nil
}

// ListUserRolesByID returns a role list of a user or an error.
func (s *Service) ListUserRolesByID(ctx context.Context, id uuid.UUID) ([]*ent.Role, error) {

	res, err := s.ReadUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.Edges.Roles, nil
}

// SubscribeUserByID returns an event list of a user or an error.
func (s *Service) SubscribeUserByID(ctx context.Context, userId uuid.UUID, eventId uuid.UUID) ([]*ent.Event, error) {

	user, err := s.ReadUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	res, err := user.Update().AddEventIDs(eventId).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, nil
}

// UnsubscribeUserByID returns an error.
func (s *Service) UnsubscribeUserByID(ctx context.Context, userId uuid.UUID, eventId uuid.UUID) error {

	user, err := s.ReadUserByID(ctx, userId)
	if err != nil {
		return err
	}
	_, err = user.Update().RemoveEventIDs(eventId).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// SetUserOnLogin returns a new or an updated user.
func (s *Service) SetUserOnLogin(ctx context.Context, user *ent.User) (*ent.User, error) {

	res, err := s.ReadUserByLogin(ctx, user.Login)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		res, err = s.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}
	} else {
		res, err = s.UpdateUserByID(ctx, res.ID, user)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func setUser(m *ent.UserMutation, u *ent.User) {

	if u.FirstName != nil {
		m.SetFirstName(*u.FirstName)
	}
	if u.LastName != nil {
		m.SetLastName(*u.LastName)
	}
	if u.DisplayName != nil {
		m.SetDisplayName(*u.DisplayName)
	}
	if u.ImagePath != nil {
		m.SetImagePath(*u.ImagePath)
	}
}
